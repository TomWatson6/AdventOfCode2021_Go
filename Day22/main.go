package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z bounds
	value   int64
}

type bounds struct {
	lower, upper int64
}

type coord struct {
	x, y, z int64
}

type section struct {
	x, y, z []int64
}

func readFile(path string) []string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	return lines
}

func getBounds(c string) bounds {
	parts := strings.Split(c, "=")
	parts = strings.Split(parts[1], "..")

	l, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}

	u, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return bounds{
		lower: l,
		upper: u,
	}
}

func parseInput(lines []string) []cube {
	var cubes []cube

	for _, line := range lines {
		if line == "" {
			continue
		}

		var ins cube

		splitLine := strings.Split(line, " ")

		if splitLine[0] == "on" {
			ins.value = 1
		} else {
			ins.value = -1
		}

		splitLine = strings.Split(splitLine[1], ",")

		ins.x = getBounds(splitLine[0])
		ins.y = getBounds(splitLine[1])
		ins.z = getBounds(splitLine[2])

		cubes = append(cubes, ins)
	}

	return cubes
}

func findIntersection(c1, c2 cube) *cube {
	if !(c1.x.lower <= c2.x.upper && c1.x.upper >= c2.x.lower) {
		return nil
	}

	if !(c1.y.lower <= c2.y.upper && c1.y.upper >= c2.y.lower) {
		return nil
	}

	if !(c1.z.lower <= c2.z.upper && c1.z.upper >= c2.z.lower) {
		return nil
	}

	x0 := int64(math.Min(float64(c1.x.lower), float64(c2.x.lower)))
	x1 := int64(math.Min(float64(c1.x.upper), float64(c2.x.upper)))
	y0 := int64(math.Min(float64(c1.y.lower), float64(c2.y.lower)))
	y1 := int64(math.Min(float64(c1.y.upper), float64(c2.y.upper)))
	z0 := int64(math.Min(float64(c1.z.lower), float64(c2.z.lower)))
	z1 := int64(math.Min(float64(c1.z.upper), float64(c2.z.upper)))

	status := c1.value * c2.value
	if c1.value == c2.value {
		status = -c1.value
	} else if c1.value == 1 && c2.value == -1 {
		status = 1
	}

	return &cube{
		x:     bounds{x0, x1},
		y:     bounds{y0, y1},
		z:     bounds{z0, z1},
		value: status,
	}
}

func getVolume(c cube) int64 {
	return (c.x.upper - c.x.lower + 1) * (c.y.upper - c.y.lower + 1) * (c.z.upper - c.z.lower + 1)
}

func main() {
	lines := readFile("input.txt")
	steps := parseInput(lines)

	var cubes []cube

	for _, step := range steps {
		var intersections []cube

		for _, cube := range cubes {
			intersection := findIntersection(step, cube)
			if intersection != nil {
				intersections = append(intersections, *intersection)
			}
		}

		cubes = append(cubes, intersections...)

		if step.value == 1 {
			cubes = append(cubes, step)
		}
	}

	res := int64(0)
	pos := int64(0)
	neg := int64(0)

	for _, cube := range cubes {
		if cube.value == 1 {
			pos += getVolume(cube) * 1
		} else {
			neg += getVolume(cube) * -1
		}
		res += getVolume(cube) * cube.value
	}

	fmt.Printf("Pos: %d, Neg: %d\n", pos, neg)

	fmt.Printf("Part 2: %d\n", res)
}

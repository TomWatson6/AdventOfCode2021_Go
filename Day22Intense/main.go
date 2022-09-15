package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type bounds struct {
	lower, upper int
}

type cube struct {
	x, y, z bounds
	on      bool
}

type coord struct {
	x, y, z int
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

	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	u, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return bounds{
		lower: l,
		upper: u + 1,
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

		ins.on = splitLine[0] == "on"

		splitLine = strings.Split(splitLine[1], ",")

		ins.x = getBounds(splitLine[0])
		ins.y = getBounds(splitLine[1])
		ins.z = getBounds(splitLine[2])

		cubes = append(cubes, ins)
	}

	return cubes
}

func main() {
	lines := readFile("simple_input.txt")
	instructions := parseInput(lines)

	var X, Y, Z []int

	for _, i := range instructions {
		X = append(X, i.x.lower)
		X = append(X, i.x.upper)
		Y = append(Y, i.y.lower)
		Y = append(Y, i.y.upper)
		Z = append(Z, i.z.lower)
		Z = append(Z, i.z.upper)
	}

	sort.Slice(X, func(i, j int) bool { return X[i] < X[j] })
	sort.Slice(Y, func(i, j int) bool { return Y[i] < Y[j] })
	sort.Slice(Z, func(i, j int) bool { return Z[i] < Z[j] })

	XI := make(map[int]int)
	YI := make(map[int]int)
	ZI := make(map[int]int)

	for i, x := range X {
		XI[x] = i
	}

	for i, y := range Y {
		YI[y] = i
	}

	for i, z := range Z {
		ZI[z] = i
	}

	grid := make(map[coord]bool)

	for i, ins := range instructions {
		fmt.Printf("Starting instruction: %d\n", i+1)

		x0 := XI[ins.x.lower]
		x1 := XI[ins.x.upper]
		y0 := YI[ins.y.lower]
		y1 := YI[ins.y.upper]
		z0 := ZI[ins.z.lower]
		z1 := ZI[ins.z.upper]

		for x := x0; x < x1; x++ {
			for y := y0; y < y1; y++ {
				for z := z0; z < z1; z++ {
					grid[coord{x, y, z}] = ins.on
				}
			}
		}
	}

	XS := len(X)
	YS := len(Y)
	ZS := len(Z)

	total := int64(0)

	for x := 0; x < XS-1; x++ {
		for y := 0; y < YS-1; y++ {
			for z := 0; z < ZS-1; z++ {
				add := int64(0)
				if on, ok := grid[coord{x, y, z}]; ok {
					if on {
						add = 1
					}
				}

				total += add * int64((X[x+1] - X[x])) * int64((Y[y+1] - Y[y])) * int64((Z[z+1] - Z[z]))
			}
		}
	}

	fmt.Printf("Part 2: %d", total)
}

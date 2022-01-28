
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Coord struct {
	X int
	Y int
}

type Line struct {
	Lower Coord
	Upper Coord
}

func (l Line) isDiag() bool {
	if l.Lower.X != l.Upper.X && l.Lower.Y != l.Upper.Y {
		return true
	} else {
		return false
	}
}

func (l *Line) sortAsc() {
	if l.Lower.X > l.Upper.X {
		temp := l.Lower.X
		l.Lower.X = l.Upper.X
		l.Upper.X = temp
	}
	if l.Lower.Y > l.Upper.Y {
		temp := l.Lower.Y
		l.Lower.Y = l.Upper.Y
		l.Upper.Y = temp
	}
}

func deleteEmpty(lines []string) []string {
	var output []string

	for _, line := range lines {
		if line != "" {
			output = append(output, line)
		}
	}

	return output
}

func toIntArray(s []string) []int {
	var output []int

	for _, char := range s {
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		output = append(output, num)
	}

	return output
}

func parseInput(inputLines []string) []Line {
	var lines []Line

	for _, line := range inputLines {
		components := strings.Split(line, " -> ")

		leftComp := strings.Split(components[0], ",")
		rightComp := strings.Split(components[1], ",")

		leftArr := toIntArray(leftComp)
		rightArr := toIntArray(rightComp)

		left := Coord{X: leftArr[0], Y: leftArr[1]}
		right := Coord{X: rightArr[0], Y: rightArr[1]}

		l := Line{Lower: left, Upper: right}
		lines = append(lines, l)
	}

	return lines
}

func printLines(lines []Line) {
	for _, line := range lines {
		fmt.Printf("%d %d -> %d %d\n", line.Lower.X, line.Lower.Y, line.Upper.X, line.Upper.Y)
	}
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)
	inputLines := strings.Split(fileString, "\n")
	inputLines = deleteEmpty(inputLines)

	lines := parseInput(inputLines)

	grid := make(map[Coord]int)

	for _, line := range lines {
		if line.isDiag() {
			continue
		}

		line.sortAsc()

		for x := line.Lower.X; x <= line.Upper.X; x++ {
			for y := line.Lower.Y; y <= line.Upper.Y; y++ {
				coord := Coord{X: x, Y: y}
				grid[coord] += 1
			}
		}
	}

	total := 0

	for _, v := range grid {
		if v > 1 {
			total += 1
		}
	}

	fmt.Printf("Part 1: %d\n", total)

	for _, line := range lines {
		if line.isDiag() {
			xStep := 1
			if line.Lower.X > line.Upper.X {
				xStep = -1
			}
			yStep := 1
			if line.Lower.Y > line.Upper.Y {
				yStep = -1
			}
			x := line.Lower.X
			y := line.Lower.Y
			for {
				coord := Coord{X: x, Y: y}
				grid[coord] += 1
				if x == line.Upper.X {
					break
				}
				x += xStep
				y += yStep
			}
		}
	}

	total = 0

	for _, v := range grid {
		if v > 1 {
			total += 1
		}
	}

	fmt.Printf("Part 2: %d\n", total)
}


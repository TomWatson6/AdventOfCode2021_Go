package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x, y int
}

func getInput(fileName string) (map[coord]rune, int, int) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fileString := string(bytes)
	lines := strings.Split(fileString, "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	grid := map[coord]rune{}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			grid[coord{x: x, y: y}] = []rune(lines[y])[x]
		}
	}

	return grid, len(lines[0]), len(lines)
}

func newCoord(x, y int) coord {
	return coord{x: x, y: y}
}

func hasNotChanged(grid map[coord]rune, evolved map[coord]rune, xs, ys int) bool {
	for k, v := range grid {
		if evolved[k] != v {
			return false
		}
	}

	return true
}

func evolve(grid map[coord]rune, xs, ys int) (map[coord]rune, bool) {
	newGrid := map[coord]rune{}
	unChanged := false

	for x := 0; x < xs; x++ {
		for y := 0; y < ys; y++ {
			if char := grid[newCoord(x, y)]; char == '>' {
				if char2 := grid[newCoord((x+1)%xs, y)]; char2 == '.' {
					newGrid[newCoord((x+1)%xs, y)] = char
					newGrid[newCoord(x, y)] = char2
				}
			}
		}
	}

	finalGrid := map[coord]rune{}

	for x := 0; x < xs; x++ {
		for y := 0; y < ys; y++ {
			if char, ok := newGrid[newCoord(x, y)]; ok {
				finalGrid[newCoord(x, y)] = char
			} else {
				finalGrid[newCoord(x, y)] = grid[newCoord(x, y)]
			}
		}
	}

	unChanged = hasNotChanged(grid, finalGrid, xs, ys)

	grid = finalGrid

	newGrid = map[coord]rune{}

	for x := 0; x < xs; x++ {
		for y := 0; y < ys; y++ {
			if char := grid[newCoord(x, y)]; char == 'v' {
				if char2 := grid[newCoord(x, (y+1)%ys)]; char2 == '.' {
					newGrid[newCoord(x, (y+1)%ys)] = char
					newGrid[newCoord(x, y)] = char2
				}
			}
		}
	}

	finalGrid = map[coord]rune{}

	for x := 0; x < xs; x++ {
		for y := 0; y < ys; y++ {
			if char, ok := newGrid[newCoord(x, y)]; ok {
				finalGrid[newCoord(x, y)] = char
			} else {
				finalGrid[newCoord(x, y)] = grid[newCoord(x, y)]
			}
		}
	}

	unChanged = unChanged && hasNotChanged(grid, finalGrid, xs, ys)

	grid = finalGrid

	return grid, unChanged
}

func main() {
	grid, xs, ys := getInput("input.txt")

	for i := 1; i < 1000; i++ {
		stopped := false
		grid, stopped = evolve(grid, xs, ys)

		if stopped {
			fmt.Printf("Part 1: %d\n", i)
			break
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	R, C int
}

type Coords []Coord

func (cs *Coords) AddIfDistinct(d Coord) {
	for _, c := range *cs {
		if c.R == d.R && c.C == d.C {
			return
		}
	}

	*cs = append(*cs, d)
}

func (cs Coords) Contains(d Coord) bool {
	for _, c := range cs {
		if c.R == d.R && c.C == d.C {
			return true
		}
	}
	return false
}

func flash(grid [][]int, coord Coord, flashed Coords) {
	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if r == c && r == 0 {
				continue
			}
			row := coord.R + r
			column := coord.C + c
			if row >= 0 && row < len(grid) && column >= 0 && column < len(grid[0]) {
				if !flashed.Contains(Coord{row, column}) {
					grid[row][column] += 1
				}
			}
		}
	}
}

func getInput(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	var output []string

	for _, line := range lines {
		if line != "" {
			line = strings.Trim(line, "\r")
			output = append(output, line)
		}
	}

	return output
}

func parseInput(lines []string) [][]int {
	var grid [][]int

	for _, line := range lines {
		chars := []rune(line)
		var row []int

		for _, char := range chars {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}

		grid = append(grid, row)
	}

	return grid
}

func main() {
	lines := getInput("input.txt")
	grid := parseInput(lines)

	steps := 100
	flashes := 0

	for i := 0; i < steps; i++ {
		var flashed Coords

		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				grid[r][c] += 1
			}
		}

		var snapshot Coords
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] > 9 {
					snapshot = append(snapshot, Coord{r, c})
				}
			}
		}

		for {
			if len(snapshot) == 0 {
				break
			}
			for _, coord := range snapshot {
				flash(grid, coord, flashed)
				flashed.AddIfDistinct(coord)
				flashes += 1
			}
			for _, coord := range snapshot {
				grid[coord.R][coord.C] = 0
			}
			snapshot = nil
			for r := 0; r < len(grid); r++ {
				for c := 0; c < len(grid[0]); c++ {
					if grid[r][c] > 9 {
						snapshot = append(snapshot, Coord{r, c})
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", flashes)

	lines = getInput("input.txt")
	grid = parseInput(lines)

	var flashed Coords
	octopuses := len(grid) * len(grid[0])
	steps = 0

	for {
		if len(flashed) == octopuses {
			break
		}

		flashed = nil

		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				grid[r][c] += 1
			}
		}

		var snapshot Coords
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] > 9 {
					snapshot = append(snapshot, Coord{r, c})
				}
			}
		}

		for {
			if len(snapshot) == 0 {
				break
			}

			for _, coord := range snapshot {
				flash(grid, coord, flashed)
				flashed.AddIfDistinct(coord)
				flashes += 1
			}
			for _, coord := range snapshot {
				grid[coord.R][coord.C] = 0
			}
			snapshot = nil
			for r := 0; r < len(grid); r++ {
				for c := 0; c < len(grid[0]); c++ {
					if grid[r][c] > 9 {
						snapshot = append(snapshot, Coord{r, c})
					}
				}
			}
		}

		steps += 1
	}

	fmt.Printf("Part 2: %d\n", steps)
}

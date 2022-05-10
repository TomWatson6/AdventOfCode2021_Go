package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Grid [][]int

type Coord struct {
	R, C int
}

type Coords []Coord

type Set []Coord

type Score map[Coord]int

func (s *Score) Initialise(width int, height int, def int) {
	*s = make(map[Coord]int)
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			coord := Coord{r, c}
			(*s)[coord] = def
		}
	}
}

func (s *Set) Add(d Coord) {
	for _, c := range *s {
		if c.R == d.R && c.C == d.C {
			return
		}
	}
	*s = append(*s, d)
}

func (s *Set) Remove(d Coord) {
	var s_ Set

	for _, c := range *s {
		if c.R != d.R || c.C != d.C {
			s_ = append(s_, c)
		}
	}

	*s = s_
}

func (s Set) Contains(d Coord) bool {
	for _, c := range s {
		if c.R == d.R && c.C == d.C {
			return true
		}
	}

	return false
}

func (cs Coords) Contains(d Coord) bool {
	for _, c := range cs {
		if c.R == d.R && c.C == d.C {
			return true
		}
	}

	return false
}

func (cs *Coords) Reverse() {
	var reversed Coords

	for i := len(*cs) - 1; i >= 0; i-- {
		reversed = append(reversed, (*cs)[i])
	}

	*cs = reversed
}

func (g *Grid) Expand(mult int) {
	var expanded Grid

	for rr := 0; rr < mult; rr++ {
		for r := 0; r < len(*g); r++ {
			var row []int
			for cc := 0; cc < mult; cc++ {
				for c := 0; c < len((*g)[0]); c++ {
					increment := rr + cc
					value := (*g)[r][c] + increment
					if value >= 10 {
						value -= 9
					}
					row = append(row, value)
				}
			}
			expanded = append(expanded, row)
		}
	}

	*g = expanded
}

func (g Grid) GetNeighbours(d Coord) Coords {
	var neighbours Coords
	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if math.Abs(float64(r)) == math.Abs(float64(c)) {
				continue
			}
			if d.R+r >= 0 && d.R+r < len(g) && d.C+c >= 0 && d.C+c < len(g[0]) {
				neighbours = append(neighbours, Coord{d.R + r, d.C + c})
			}
		}
	}

	return neighbours
}

func (g Grid) GetHeuristicValue(start Coord, dest Coord) int {
	r := dest.R - start.R
	c := dest.C - start.C
	return r + c
}

func (g Grid) ReconstructPath(cameFrom map[Coord]Coord, c Coord) (Coords, int) {
	var totalPath Coords
	totalPath = append(totalPath, c)
	totalDistance := g[c.R][c.C]

	for {
		if _, exists := cameFrom[c]; !exists {
			break
		}
		c = cameFrom[c]
		totalDistance += g[c.R][c.C]
		totalPath = append(totalPath, c)
	}
	totalPath.Reverse()

	return totalPath, totalDistance
}

func (g Grid) FindPath(start Coord, goal Coord) (Coords, int) {
	var openSet Set
	openSet.Add(start)

	cameFrom := make(map[Coord]Coord)

	width := len(g[0])
	height := len(g)

	var gScore Score
	gScore.Initialise(width, height, 999999)
	gScore[start] = 0

	var fScore Score
	fScore.Initialise(width, height, 999999)
	fScore[start] = g.GetHeuristicValue(start, goal)

	for {
		if len(openSet) == 0 {
			break
		}
		var current Coord
		minScore := 999999
		for _, c := range openSet {
			if fScore[c] < minScore {
				minScore = fScore[c]
				current = c
			}
		}
		if current == goal {
			return g.ReconstructPath(cameFrom, current)
		}
		openSet.Remove(current)
		neighbours := g.GetNeighbours(current)
		for _, n := range neighbours {
			tentativeGScore := gScore[current] + g[n.R][n.C]
			if tentativeGScore < gScore[n] {
				cameFrom[n] = current
				gScore[n] = tentativeGScore
				fScore[n] = tentativeGScore + g.GetHeuristicValue(n, goal)
				if !openSet.Contains(n) {
					openSet.Add(n)
				}
			}
		}
	}
	return Coords{}, -1
}

func getInput(fileName string) Grid {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	var rows []string

	for _, line := range lines {
		if line != "" {
			line = strings.Trim(line, "\r")
			rows = append(rows, line)
		}
	}

	var grid Grid

	for _, row := range rows {
		var nums []int

		for _, char := range []rune(row) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		grid = append(grid, nums)
	}

	return grid
}

func main() {
	grid := getInput("input.txt")
	_, distance := grid.FindPath(Coord{0, 0}, Coord{len(grid) - 1, len(grid[0]) - 1})
	fmt.Printf("Part 1: %d\n", distance-grid[0][0])

	grid.Expand(5)
	_, distance = grid.FindPath(Coord{0, 0}, Coord{len(grid) - 1, len(grid[0]) - 1})
	fmt.Printf("Part 2: %d\n", distance-grid[0][0])
}

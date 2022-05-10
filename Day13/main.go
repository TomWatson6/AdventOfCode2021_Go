package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type Coords []Coord

type Fold struct {
	IsX      bool
	Position int
}

func (cs *Coords) Add(c Coord) {
	for _, coord := range *cs {
		if coord.X == c.X && coord.Y == c.Y {
			return
		}
	}
	*cs = append(*cs, c)
}

func (cs Coords) Print() {
	xMin := 99999999
	yMin := 99999999
	xMax := 0
	yMax := 0

	for _, c := range cs {
		if c.X < xMin {
			xMin = c.X
		}
		if c.X > xMax {
			xMax = c.X
		}
		if c.Y < yMin {
			yMin = c.Y
		}
		if c.Y > yMax {
			yMax = c.Y
		}
	}

	coordMap := make(map[Coord]bool)

	for _, c := range cs {
		coordMap[c] = true
	}

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if coordMap[Coord{x, y}] {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", ".")
			}
		}
		fmt.Println()
	}
}

func (cs *Coords) Fold(f Fold) {
	var next Coords

	for i := 0; i < len(*cs); i++ {
		if f.IsX {
			if (*cs)[i].X > f.Position {
				x := f.Position - ((*cs)[i].X - f.Position)
				next.Add(Coord{x, (*cs)[i].Y})
			} else {
				next.Add((*cs)[i])
			}
		} else {
			if (*cs)[i].Y > f.Position {
				y := f.Position - ((*cs)[i].Y - f.Position)
				next.Add(Coord{(*cs)[i].X, y})
			} else {
				next.Add((*cs)[i])
			}
		}
	}

	*cs = next
}

func getInput(fileName string) ([]Fold, Coords) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	components := strings.Split(fileString, "\n")

	var lines []string

	for _, line := range components {
		if line != "" {
			line = strings.Trim(line, "\r")
			lines = append(lines, line)
		}
	}

	var coords Coords
	var folds []Fold

	for _, line := range lines {
		if line == "" {
			continue
		}
		chars := []rune(line)
		if string(chars[:4]) == "fold" {
			parts := strings.Split(line, " ")
			fold := strings.Split(parts[len(parts)-1], "=")
			amount, err := strconv.Atoi(fold[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Fold{fold[0] == "x", amount})
		} else {
			parts := strings.Split(line, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			coords = append(coords, Coord{x, y})
		}
	}

	return folds, coords
}

func main() {
	folds, coords := getInput("input.txt")

	for i := 0; i < len(folds); i++ {
		coords.Fold(folds[i])

		if i == 0 {
			fmt.Printf("Part 1: %d\n", len(coords))
		}
	}

	coords.Print()
}

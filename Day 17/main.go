package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type CoordSet []Coord

func (s *CoordSet) Add(c Coord) {
	for _, coord := range *s {
		if coord.X == c.X && coord.Y == c.Y {
			return
		}
	}

	*s = append(*s, c)
}

func (s *CoordSet) Remove(c Coord) {
	var final CoordSet

	for _, coord := range *s {
		if coord.X != c.X && coord.Y != c.Y {
			final = append(final, coord)
		}
	}

	*s = final
}

type Bound struct {
	Lower, Upper int
}

type Area struct {
	X, Y Bound
}

type Set []int

func (s Set) Add(n int) Set {
	for _, v := range s {
		if v == n {
			return s
		}
	}
	s = append(s, n)

	return s
}

func (s *Set) Remove(n int) {
	var final []int
	for _, v := range *s {
		if v != n {
			final = append(final, v)
		}
	}
	*s = final
}

func GetArea(fileName string) Area {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(bytes)
	parts := strings.Split(fileString, " ")
	x := strings.Trim(parts[2], ",")
	y := parts[3]
	splitX := strings.Split(strings.Trim(x, "x="), "..")
	splitY := strings.Split(strings.Trim(y, "y="), "..")

	fmt.Printf("X: %v, Y: %v\n", splitX, splitY)

	xLower, _ := strconv.Atoi(splitX[0])
	xUpper, _ := strconv.Atoi(splitX[1])
	yLower, _ := strconv.Atoi(splitY[0])
	yUpper, _ := strconv.Atoi(splitY[1])

	xBound := Bound{xLower, xUpper}
	yBound := Bound{yLower, yUpper}

	area := Area{xBound, yBound}

	return area
}

func Calculate(number int) int {
	if number <= 1 {
		return 1
	} else {
		return number + Calculate(number-1)
	}
}

func GetX(initial int, time int) int {
	x := 0
	for {
		if time <= 0 {
			break
		}
		x += initial

		if initial > 0 {
			initial -= 1
		} else if initial < 0 {
			initial += 1
		}
		time -= 1
	}

	return x
}

func GetY(initial int, time int) int {
	y := 0

	for {
		if time <= 0 {
			break
		}

		y += initial

		initial -= 1
		time -= 1
	}

	return y
}

func main() {
	targetArea := GetArea("input.txt")
	fmt.Printf("Target Area: %v\n", targetArea)

	xStep := 1

	medianX := (targetArea.X.Lower + targetArea.X.Upper) / 2

	if medianX < 0 {
		xStep = -1
	}

	xValue := 0

	for x := 0; x != xStep*1000; x += xStep {
		xFinal := Calculate(x)
		if xFinal >= targetArea.X.Lower && xValue == 0 {
			xValue = x
			break
		}
	}

	yStart := int(math.Abs(float64(targetArea.Y.Lower))) - 1
	yMax := Calculate(yStart)

	fmt.Printf("Part 1: %d\n", yMax)

	xTimes := make(map[int]Set)
	yTimes := make(map[int]Set)

	for t := 1; t < (yStart+1)*3; t++ {
		for x := xValue; x < targetArea.X.Upper+1; x++ {
			hit := GetX(x, t)
			if targetArea.X.Lower <= hit && hit <= targetArea.X.Upper {
				if xTimes[t] == nil {
					xTimes[t] = Set{}
				}
				xTimes[t] = xTimes[t].Add(x)
			}
		}

		for y := targetArea.Y.Lower; y < yStart+1; y++ {
			hit := GetY(y, t)
			if targetArea.Y.Lower <= hit && hit <= targetArea.Y.Upper {
				if yTimes[t] == nil {
					yTimes[t] = Set{}
				}
				yTimes[t] = yTimes[t].Add(y)
			}
		}
	}

	var points CoordSet
	hits := 0

	for t := 1; t < (yStart+1)*3; t++ {
		for _, sx := range xTimes[t] {
			for _, sy := range yTimes[t] {
				points.Add(Coord{sx, sy})
			}
		}
	}

	hits = len(points)

	fmt.Printf("Part 2: %d\n", hits)
}

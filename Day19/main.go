package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type Beacon struct {
	X, Y, Z int
}

var rotations = []Beacon{
	{0, 0, 0},
	{0, 0, 1},
	{0, 0, 2},
	{0, 0, 3},
	{1, 0, 0},
	{1, 0, 1},
	{1, 0, 2},
	{1, 0, 3},
	{2, 0, 0},
	{2, 0, 1},
	{2, 0, 2},
	{2, 0, 3},
	{3, 0, 0},
	{3, 0, 1},
	{3, 0, 2},
	{3, 0, 3},
	{0, 1, 0},
	{0, 1, 1},
	{0, 1, 2},
	{0, 1, 3},
	{0, 3, 0},
	{0, 3, 1},
	{0, 3, 2},
	{0, 3, 3},
}

func (b Beacon) Rotate(x, y, z int) Beacon {
	for i := 0; i < x; i++ {
		b.Z, b.Y = -b.Y, b.Z
	}
	for i := 0; i < y; i++ {
		b.Z, b.X = -b.X, b.Z
	}
	for i := 0; i < z; i++ {
		b.Y, b.X = -b.X, b.Y
	}
	return b
}

func GetRelativePosition(b1, b2 Beacon) Beacon {
	return Beacon{b1.X - b2.X, b1.Y - b2.Y, b1.Z - b2.Z}
}

type Beacons []Beacon

func (bs Beacons) Rotate(x, y, z int) Beacons {
	var b2 Beacons
	for _, b := range bs {
		b2 = append(b2, b.Rotate(x, y, z))
	}
	return b2
}

func (bs Beacons) GetMatches(bs2 Beacons) (Beacon, int) {
	relPos := make(map[Beacon]int)

	for _, b := range bs {
		for _, b2 := range bs2 {
			relPos[GetRelativePosition(b, b2)]++
		}
	}

	highest := 0
	highestBeacon := Beacon{}

	for k, v := range relPos {
		if v > highest {
			highest = v
			highestBeacon = k
		}
	}

	return highestBeacon, highest
}

type Scanner struct {
	Id        int
	Beacons   Beacons
	Rotations []Beacons
}

func (s *Scanner) GenerateRotations() {
	s.Rotations = nil
	for _, r := range rotations {
		s.Rotations = append(s.Rotations, s.Beacons.Rotate(r.X, r.Y, r.Z))
	}
}

func NewScanner(id int, beacons Beacons) Scanner {
	var scanner Scanner
	scanner.Id = id
	scanner.Beacons = beacons
	scanner.GenerateRotations()

	return scanner
}

func getInput(fileName string) []Scanner {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(bytes)
	scannerParts := strings.Split(fileString, "\n\n")
	if len(scannerParts) == 1 {
		scannerParts = strings.Split(fileString, "\r\n\r\n")
	}

	var scanners []Scanner

	for _, scannerPart := range scannerParts {
		var beacons Beacons
		parts := strings.Split(scannerPart, "\r\n")
		if len(parts) == 1 {
			parts = strings.Split(scannerPart, "\n")
		}
		idString := parts[0]
		idString = strings.Replace(idString, "--- scanner ", "", -1)
		idString = strings.Replace(idString, " ---", "", -1)
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		for _, line := range parts[1:] {
			if line == "" {
				continue
			}
			coordStrings := strings.Split(line, ",")
			x, err := strconv.Atoi(coordStrings[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(coordStrings[1])
			if err != nil {
				panic(err)
			}
			z, err := strconv.Atoi(coordStrings[2])
			if err != nil {
				panic(err)
			}
			beacons = append(beacons, Beacon{x, y, z})
		}
		scanners = append(scanners, NewScanner(id, beacons))
	}
	return scanners
}

type ScannerTree struct {
	Id            int
	Position      Beacon
	Scanner       Scanner
	ScannersTried map[int]bool
	Children      map[int]*ScannerTree
}

func (s ScannerTree) Print(level int) {
	output := ""
	for i := 0; i < level; i++ {
		if i == level-1 {
			output += "|"
		} else {
			output += " "
		}
	}
	idString := strconv.Itoa(s.Id)
	output += idString
	fmt.Printf("%s\n", output)
	for _, child := range s.Children {
		step := 3
		if len(idString) == 1 {
			step = 2
		}
		child.Print(level + step)
	}
}

func (s *ScannerTree) AddChild(child *ScannerTree) bool {
	if ok := s.ScannersTried[child.Id]; !ok {
		// ctx, cancel := context.WithCancel(context.Background())
		for _, r := range child.Scanner.Rotations {
			if relPos, count := s.Scanner.Beacons.GetMatches(r); count >= 12 {
				child.Position = relPos
				child.Scanner.Beacons = r
				child.Scanner.GenerateRotations()
				s.Children[child.Id] = child
				return true
			}
		}
	}

	for _, c := range s.Children {
		if ok := c.AddChild(child); ok {
			return true
		}
	}

	s.ScannersTried[child.Id] = true
	return false
}

func (s ScannerTree) GetBeacons() (map[Beacon]bool, int) {
	beacons := make(map[Beacon]bool)

	for _, b := range s.Scanner.Beacons {
		beacons[b] = true
	}

	for _, c := range s.Children {
		cBeacons, _ := c.GetBeacons()
		for k := range cBeacons {
			beacons[Beacon{k.X + c.Position.X, k.Y + c.Position.Y, k.Z + c.Position.Z}] = true
		}
	}

	count := 0

	for i := 0; i < len(beacons); i++ {
		count += 1
	}

	return beacons, count
}

func (s ScannerTree) GetAbsolutePositions(p Beacon) Beacons {
	pos := Beacon{
		X: s.Position.X + p.X,
		Y: s.Position.Y + p.Y,
		Z: s.Position.Z + p.Z,
	}

	var positions Beacons
	positions = append(positions, pos)

	for _, c := range s.Children {
		cPositions := c.GetAbsolutePositions(pos)
		positions = append(positions, cPositions...)
	}

	return positions
}

func main() {
	scanners := getInput("Day19/input.txt")

	start := time.Now()
	if len(scanners) > 0 {
		scannerTree := ScannerTree{
			Id:            scanners[0].Id,
			Scanner:       scanners[0],
			ScannersTried: make(map[int]bool),
			Children:      make(map[int]*ScannerTree),
		}

		var trees []*ScannerTree

		for _, scanner := range scanners[1:] {
			trees = append(trees, &ScannerTree{
				Id:            scanner.Id,
				Scanner:       scanner,
				ScannersTried: make(map[int]bool),
				Children:      make(map[int]*ScannerTree),
			})
		}

		for {
			if len(trees) == 0 {
				break
			}

			var toRemove []int

			for i, scanner := range trees {
				if ok := scannerTree.AddChild(scanner); ok {
					toRemove = append(toRemove, i)
				}
			}

			for i := len(toRemove) - 1; i >= 0; i-- {
				trees = append(trees[:toRemove[i]], trees[toRemove[i]+1:]...)
			}
		}

		_, count := scannerTree.GetBeacons()

		t := time.Since(start)

		fmt.Printf("Part 1: %d - %dms\n", count, t.Milliseconds())

		start = time.Now()

		positions := scannerTree.GetAbsolutePositions(Beacon{0, 0, 0})

		largest := 0

		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				x0, y0, z0 := positions[i].X, positions[i].Y, positions[i].Z
				x1, y1, z1 := positions[j].X, positions[j].Y, positions[j].Z
				distance := math.Abs(float64(x0-x1)) + math.Abs(float64(y0-y1)) + math.Abs(float64(z0-z1))
				if int(distance) > largest {
					largest = int(distance)
				}
			}
		}

		t = time.Since(start)

		fmt.Printf("Part 2: %d - %dms\n", largest, t.Milliseconds())
	} else {
		panic("No scanners found in file")
	}
}

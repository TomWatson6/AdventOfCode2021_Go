package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type dungeon map[coord]tile

type coord struct {
	x, y int
}

type tile byte

const (
	empty tile = iota
	available
	wall
	A
	B
	C
	D
)

var cost = map[tile]int{
	A: 1,
	B: 10,
	C: 100,
	D: 1000,
}

var dest = map[tile]int{
	A: 3,
	B: 5,
	C: 7,
	D: 9,
}

var corrY = 0 // Needs setting when reading input

func (d dungeon) distance(c coord) int {
	t, ok := d[c]
	if !ok {
		panic("trying to move tile that doesn't exist")
	}

	if c.x == dest[t] {
		moves := 0
		next := coord{c.x, c.y + 1}
		for d[next] != wall || d[next] != t {
			next = coord{next.x, next.y + 1}
			moves += 1
		}
	}

	return -1
}

// Make a tree that represents the map, i.e. left / right / contents (when empty unblocked, when occupied, blocks path between adjacent nodes)
// 	\
// 	 \
// 	D \
// C   \
// 	  A \
// 	 A   \
// 	    D \
// 	   B   \
// 	      C \
// 		 B   \

// #############
// #...........#
// ###D#A#D#C###
//   #C#A#B#B#
//   #########

func getInput(fileName string) (*dungeon, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fileString := string(bytes)
	lines := strings.Split(fileString, "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	d := make(dungeon)

	for y := range lines {
		for x := range lines[y] {
			switch lines[y][x] {
			case 'A':
				d[coord{x, y}] = A
			case 'B':
				d[coord{x, y}] = B
			case 'C':
				d[coord{x, y}] = C
			case 'D':
				d[coord{x, y}] = D
			case '.':
				corrY = y
				d[coord{x, y}] = available
			case '#':
				d[coord{x, y}] = wall
			case ' ':
				d[coord{x, y}] = empty
			}
		}
	}

	return &d, nil
}

func (d dungeon) getAvailableMoves(c coord) []coord {
	moves := []int{-1, 1}
	var tiles []coord

	for x := range moves {
		for y := range moves {
			c0 := coord{x: c.x + x, y: c.y + y}
			if d[c0] == empty {
				tiles = append(tiles, c0)
			}
		}
	}

	return tiles
}

func (d dungeon) findPath(start, dest coord) []coord {
	// A* path planning algorithm
	return []coord{}
}

func (d dungeon) calculateCost(path []coord) int {
	amp := d[path[0]]

	c := cost[amp] * (len(path) - 1)

	return c
}

// remaining steps are to find the columns to prioritise - the emptier ones need buffering to move correct into it

func main() {
	d, err := getInput("simple_input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", d)
}

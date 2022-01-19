
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")
	lines = delete_empty(lines)

	x := 0
	y := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		instruction := components[0]
		magnitude, err := strconv.Atoi(components[1])
		if err != nil {
			panic(err)
		}

		if instruction == "forward" {
			x += magnitude
		}
		if instruction == "down" {
			y += magnitude
		}
		if instruction == "up" {
			y -= magnitude
		}
	}

	fmt.Printf("Part 1: %d\n", x * y)

	x = 0
	y = 0
	aim := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		instruction := components[0]
		magnitude, err := strconv.Atoi(components[1])
		if err != nil {
			panic(err)
		}

		if instruction == "forward" {
			x += magnitude
			y += magnitude * aim
		}
		if instruction == "down" {
			aim += magnitude
		}
		if instruction == "up" {
			aim -= magnitude
		}
	}
	fmt.Printf("Part 2: %d\n", x * y)
}


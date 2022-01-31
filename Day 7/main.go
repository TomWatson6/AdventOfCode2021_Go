package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func calculateFuelCosts(positions []int, target int) uint64 {
	total := uint64(0)

	for _, pos := range positions {
		total += uint64(math.Abs(float64(pos - target)))
	}

	return total
}

func calculateFuelCosts2(positions []int, target int) uint64 {
	total := uint64(0)

	for _, pos := range positions {
		diff := uint64(math.Abs(float64(pos - target)))
		total += (diff * (diff + 1)) / 2
	}

	return total
}

func toIntArray(input []string) []int {
	var output []int

	for _, line := range input {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		output = append(output, i)
	}

	return output
}

func getInput(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, ",")
	output := toIntArray(lines)

	return output
}

func main() {
	positions := getInput("input.txt")

	sort.Ints(positions)
	median := positions[len(positions)/2]
	fuelCosts := calculateFuelCosts(positions, median)

	fmt.Printf("Part 1: %d\n", fuelCosts)

	min := 99999999
	max := 0

	for _, pos := range positions {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	smallest := uint64(9999999999999)

	for i := min; i <= max; i++ {
		fuelCost := calculateFuelCosts2(positions, i)
		if fuelCost < smallest {
			smallest = fuelCost
		}
	}

	fmt.Printf("Part 2: %d\n", smallest)
}

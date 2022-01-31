package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func toIntArray(lines []string) []int {
	var output []int

	for _, line := range lines {
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

func parseFish(numbers []int) map[int]uint64 {
	fish := make(map[int]uint64)

	for _, number := range numbers {
		fish[number]++
	}

	return fish
}

func evolve(fish map[int]uint64, days int) uint64 {
	new_fish := uint64(0)

	for i := 0; i < days; i++ {
		new_fish = fish[0]
		for x := 0; x < 8; x++ {
			fish[x] = fish[x+1]
		}
		fish[6] += new_fish
		fish[8] = new_fish
	}

	sum := uint64(0)

	for _, v := range fish {
		sum += v
	}

	return sum
}

func main() {
	numbers := getInput("input.txt")
	fish := parseFish(numbers)

	p1 := evolve(fish, 80)
	fmt.Printf("Part 1: %d\n", p1)

	fish = parseFish(numbers)
	p2 := evolve(fish, 256)
	fmt.Printf("Part 2: %d\n", p2)
}

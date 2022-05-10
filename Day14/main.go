package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func evolve(pairs map[string]uint64, rules map[string]string) map[string]uint64 {
	evolved := make(map[string]uint64)

	for k, v := range pairs {
		if v == 0 {
			continue
		}
		insertion := rules[k]
		if len(insertion) == 0 {
			continue
		}
		first := string([]rune(k)[0])
		last := string([]rune(k)[1])
		evolved[first+insertion] += v
		evolved[insertion+last] += v
	}

	return evolved
}

func countLetters(pairs map[string]uint64, firstLetter rune, lastLetter rune) uint64 {
	totals := make(map[rune]uint64)

	for k, v := range pairs {
		for _, letter := range []rune(k) {
			totals[letter] += v
		}
	}

	for k := range totals {
		if k == firstLetter || k == lastLetter {
			totals[k] += uint64(1)
		}
	}

	for k, v := range totals {
		fmt.Printf("%s: %d\n", string(k), v)
	}

	return getDifference(totals, firstLetter, lastLetter)
}

func getDifference(totals map[rune]uint64, firstLetter rune, lastLetter rune) uint64 {
	min := uint64(99999999)
	max := uint64(0)

	for _, v := range totals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return (max - min) / 2
}

func getInput(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	var input []string

	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line != "" {
			input = append(input, line)
		}
	}

	return input
}

func parseInput(lines []string) (map[string]uint64, map[string]string) {
	template := []rune(lines[0])
	pairs := make(map[string]uint64)

	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		pairs[pair] += uint64(1)
	}

	lines = lines[2:]
	rules := make(map[string]string)

	for _, line := range lines {
		split := strings.Split(line, " -> ")
		left := split[0]
		right := split[1]
		rules[left] = right
	}

	return pairs, rules
}

func main() {
	lines := getInput("simple_input.txt")
	pairs, rules := parseInput(lines)
	letters := []rune(lines[0])
	firstLetter := letters[0]
	lastLetter := letters[len(letters)-1]

	for i := 0; i < 11; i++ {
		if i == 10 {
			fmt.Printf("Part 1: %d\n", countLetters(pairs, firstLetter, lastLetter))
		}
		pairs = evolve(pairs, rules)
	}

	fmt.Printf("Part 2: %d\n", countLetters(pairs, firstLetter, lastLetter))
}

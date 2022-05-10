package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Display struct {
	Input, Output []string
	Patterns      map[rune][]int
	StringMap     map[string]int
}

func (d Display) GetNumAppearances(nums []int) int {
	total := 0

	for _, item := range d.Output {
		itemRunes := []rune(item)
		var itemStrings []string
		for _, char := range itemRunes {
			itemStrings = append(itemStrings, string(char))
		}
		sort.Strings(itemStrings)
		sorted := ""
		for _, char := range itemStrings {
			sorted += char
		}
		itemNum := d.StringMap[sorted]
		for _, num := range nums {
			if itemNum == num {
				total += 1
				break
			}
		}
	}

	return total
}

func (d *Display) InitStringMap(defaultPatterns map[rune][]int, defaultStringMap map[string]int) {
	charMap := make(map[rune]rune)

	for k, v := range d.Patterns {
		for k1, v1 := range defaultPatterns {
			match := true

			if len(v) != len(v1) {
				continue
			}

			for i := range v {
				if v[i] != v1[i] {
					match = false
				}
			}

			if match {
				charMap[k] = k1
				break
			}
		}
	}

	// For each d.Input, replace characters using char map, find number from default string map, map original string to number found
	for _, s := range d.Input {
		chars := []rune(s)

		orderedInputString := orderString(s)

		for i := 0; i < len(chars); i++ {
			chars[i] = charMap[chars[i]]
		}
		var charStrings []string
		for _, c := range chars {
			charStrings = append(charStrings, string(c))
		}
		sort.Strings(charStrings)
		mapped := ""
		for _, t := range charStrings {
			mapped += string([]rune(t)[0])
		}
		d.StringMap[orderedInputString] = defaultStringMap[mapped]
	}
}

func (d Display) GetOutput() int {
	output := ""

	for _, value := range d.Output {
		orderedValue := orderString(value)
		v := strconv.Itoa(d.StringMap[orderedValue])
		output += v
	}

	outputValue, err := strconv.Atoi(output)
	if err != nil {
		panic(err)
	}

	return outputValue
}

func getPatterns(sequences []string) map[rune][]int {
	patterns := make(map[rune][]int)

	for c := 97; c < 104; c++ {
		char := []rune(string(c))[0]
		var setSizes []int
		for _, v := range sequences {
			if contains(v, char) {
				setSizes = append(setSizes, len(v))
			}
		}
		sort.Ints(setSizes)
		patterns[char] = setSizes
	}

	return patterns
}

func orderString(s string) string {
	chars := []rune(s)
	var str []string

	for _, c := range chars {
		str = append(str, string(c))
	}

	sort.Strings(str)
	joined := strings.Join(str, "")

	return joined
}

func getDefaultNumbers() []string {
	var numbers = []string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}

	return numbers
}

func getStringMap() map[string]int {
	stringMap := make(map[string]int)

	stringMap["abcefg"] = 0
	stringMap["cf"] = 1
	stringMap["acdeg"] = 2
	stringMap["acdfg"] = 3
	stringMap["bcdf"] = 4
	stringMap["abdfg"] = 5
	stringMap["abdefg"] = 6
	stringMap["acf"] = 7
	stringMap["abcdefg"] = 8
	stringMap["abcdfg"] = 9

	return stringMap
}

func getInput(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	return lines
}

func contains(array string, value rune) bool {
	for _, char := range []rune(array) {
		if char == value {
			return true
		}
	}
	return false
}

func parseInput(lines []string) []Display {
	var displays []Display

	for _, line := range lines {
		parts := strings.Split(line, "|")

		for i := 0; i < len(parts); i++ {
			parts[i] = strings.Trim(parts[i], " ")
		}

		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		patterns := getPatterns(input)
		displays = append(displays, Display{input, output, patterns, make(map[string]int)})
	}

	return displays
}

func main() {
	lines := getInput("input.txt")
	displays := parseInput(lines)
	defaultNumbers := getDefaultNumbers()
	stringMap := getStringMap()
	defaultPatterns := getPatterns(defaultNumbers)

	for i := 0; i < len(displays); i++ {
		displays[i].InitStringMap(defaultPatterns, stringMap)
	}

	var p1Nums = []int{1, 4, 7, 8}

	total := 0

	for _, display := range displays {
		total += display.GetNumAppearances(p1Nums)
	}

	fmt.Printf("Part 1: %d\n", total)

	total = 0

	for _, display := range displays {
		total += display.GetOutput()
	}

	fmt.Printf("Part 2: %d\n", total)
}

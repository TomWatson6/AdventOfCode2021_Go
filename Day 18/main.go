package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
)

var ()

func Contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}

	return false
}

func Remove(arr []string, start int, end int) []string {
	var output []string

	for i := range arr {
		if i < start || i >= end {
			output = append(output, arr[i])
		}
	}

	return output
}

func Insert(arr []string, insertion []string, index int) ([]string, error) {
	if index > len(arr)-1 {
		return nil, fmt.Errorf("Index out of range for insertion: %d\n", index)
	}

	var output []string

	for i := 0; i < index; i++ {
		output = append(output, arr[i])
	}
	for _, item := range insertion {
		output = append(output, item)
	}
	for i := index; i < len(arr); i++ {
		output = append(output, arr[i])
	}

	return output, nil
}

func Simplify(pair []string) []string {
	lastNumber := -1
	depth := 0
	i := 0
	nonNumbers := []string{
		"[",
		"]",
		",",
	}

	for {
		if i >= len(pair) {
			break
		}
		if pair[i] == "[" {
			depth += 1
		} else if pair[i] == "]" {
			depth -= 1
		} else if pair[i] != "," {
			if depth > 4 && pair[i+1] != "]" && pair[i+3] == "]" {
				// EXPLODE ONLY IF FORMAT [A, B]
				if lastNumber != -1 {
					firstNumber, err := strconv.Atoi(pair[lastNumber])
					if err != nil {
						panic(err)
					}
					secondNumber, err := strconv.Atoi(pair[i])
					if err != nil {
						panic(err)
					}
					pair[lastNumber] = strconv.Itoa(firstNumber + secondNumber)
				}
				i += 2
				end := i + 1
				for {
					if end >= len(pair) {
						break
					}
					if !Contains(nonNumbers, pair[end]) {
						firstNumber, err := strconv.Atoi(pair[end])
						if err != nil {
							panic(err)
						}
						secondNumber, err := strconv.Atoi(pair[i])
						if err != nil {
							panic(err)
						}
						pair[end] = strconv.Itoa(firstNumber + secondNumber)
						break
					}
					end += 1
				}
				pair[i-3] = "0"
				pair = Remove(pair, i-2, i+2)
				return pair
			} else {
				lastNumber = i
			}
		}
		i += 1
	}
	i = 0
	depth = 0
	// Check for splits
	for {
		if i >= len(pair) {
			break
		}
		if !Contains(nonNumbers, pair[i]) {
			if len(pair[i]) > 1 {
				num, err := strconv.Atoi(pair[i])
				if err != nil {
					panic(err)
				}
				left := int(math.Floor(float64(num) / 2))
				right := int(math.Ceil(float64(num) / 2))
				insertion := []string{
					"[",
					strconv.Itoa(left),
					",",
					strconv.Itoa(right),
					"]",
				}
				pair = Remove(pair, i, i+1)
				pair, err = Insert(pair, insertion, i)
				if err != nil {
					panic(err)
				}
				return pair
			}
		}
		i += 1
	}
	return pair
}

func FindMagnitude(pairs []string) int {
	var totals []int
	pos := 0
	for {
		if pos >= len(pairs) {
			break
		}
		if pairs[pos] == "[" {
			nestingLevel := 1
			end := pos + 1
			for {
				if nestingLevel <= 0 {
					break
				}
				if pairs[end] == "[" {
					nestingLevel += 1
				} else if pairs[end] == "]" {
					nestingLevel -= 1
				}
				end += 1
			}
			totals = append(totals, FindMagnitude(pairs[pos+1:end-1]))
			pos = end
		} else if pairs[pos] != "," {
			number, err := strconv.Atoi(pairs[pos])
			if err != nil {
				panic(err)
			}
			totals = append(totals, number)
		}
		pos += 1
	}

	if len(totals) > 1 {
		return (totals[0] * 3) + (totals[1] * 2)
	} else {
		return totals[0]
	}
}

func GetInput(fileName string) [][]string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(bytes)
	lines := strings.Split(fileString, "\n")

	var output []string

	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line != "" {
			output = append(output, line)
		}
	}

	var strArr [][]string

	for _, line := range output {
		runes := []rune(line)
		var toAdd []string
		for _, r := range runes {
			toAdd = append(toAdd, string(r))
		}
		strArr = append(strArr, toAdd)
	}

	return strArr
}

func Equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Join(a []string, b []string) []string {
	var line []string

	line = append(line, "[")

	for _, item := range a {
		line = append(line, item)
	}

	line = append(line, ",")

	for _, item := range b {
		line = append(line, item)
	}

	line = append(line, "]")

	return line
}

func GetMagnitude(a []string, b []string) int {
	line := Join(a, b)

	for {
		var temp []string

		for _, item := range line {
			temp = append(temp, item)
		}

		line = Simplify(line)

		if Equals(line, temp) {
			break
		}
	}

	magnitude := FindMagnitude(line)

	return magnitude
}

func main() {
	lines := GetInput("input.txt")

	for i := 0; i < len(lines)-1; i++ {
		line := Join(lines[i], lines[i+1])

		for {
			var temp []string

			for _, item := range line {
				temp = append(temp, item)
			}

			line = Simplify(line)

			if Equals(temp, line) {
				break
			}
		}

		lines[i+1] = line
	}

	magnitude := FindMagnitude(lines[len(lines)-1])

	fmt.Printf("Part 1: %d\n", magnitude)

	lines = GetInput("input.txt")

	var magnitudes []int
	var wg = &sync.WaitGroup{}

	for i1 := 0; i1 < len(lines)-1; i1++ {
		for i2 := i1; i2 < len(lines); i2++ {
			wg.Add(2)
			go func(i1 int, i2 int) {
				magnitudes = append(magnitudes, GetMagnitude(lines[i1], lines[i2]))
				wg.Done()
				magnitudes = append(magnitudes, GetMagnitude(lines[i2], lines[i1]))
				wg.Done()
			}(i1, i2)
		}
	}

	wg.Wait()

	max := 0

	for _, magnitude := range magnitudes {
		if magnitude > max {
			max = magnitude
		}
	}

	fmt.Printf("Part 2: %d\n", max)
}

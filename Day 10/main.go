package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Stack []rune

func (s *Stack) Push(char rune) {
	*s = append(*s, char)
	// s = &arr
}

func (s *Stack) Pop() rune {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func getInput(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")
	var output []string

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	for _, line := range lines {
		if line != "" {
			output = append(output, line)
		}
	}

	return output
}

func main() {
	scoring := make(map[rune]int)
	scoring[')'] = 3
	scoring[']'] = 57
	scoring['}'] = 1197
	scoring['>'] = 25137

	bracketMap := make(map[rune]rune)
	bracketMap['('] = ')'
	bracketMap['['] = ']'
	bracketMap['{'] = '}'
	bracketMap['<'] = '>'

	lines := getInput("input.txt")
	totalCorr := 0
	score := 0
	var incomplete []string

	for _, line := range lines {
		var stack Stack
		corrupted := false
		chars := []rune(line)
		openBrackets := "([{<"

		for _, c := range chars {
			if strings.Contains(openBrackets, string(c)) {
				stack.Push(c)
			} else if c == bracketMap[stack[len(stack)-1]] {
				stack.Pop()
			} else {
				corrupted = true
				score += scoring[c]
				break
			}
		}
		if len(stack) > 0 && !corrupted {
			incomplete = append(incomplete, line)
		}
		if corrupted {
			totalCorr += 1
		}
	}

	fmt.Printf("Part 1: %d\n", score)

	charValues := make(map[rune]int)
	charValues[')'] = 1
	charValues[']'] = 2
	charValues['}'] = 3
	charValues['>'] = 4

	var totals []int

	for _, line := range incomplete {
		total := 0
		var stack Stack
		chars := []rune(line)
		openBrackets := "([{<"

		for _, c := range chars {
			if strings.Contains(openBrackets, string(c)) {
				stack.Push(c)
			} else if c == bracketMap[stack[len(stack)-1]] {
				stack.Pop()
			}
		}

		var reversed Stack
		length := len(stack)
		for i := 0; i < length; i++ {
			reversed.Push(stack.Pop())
		}

		for _, c := range reversed {
			char := bracketMap[c]
			total *= 5
			total += charValues[char]
		}

		totals = append(totals, total)
	}

	sort.Ints(totals)
	fmt.Printf("Part 2: %d\n", totals[len(totals)/2])
}

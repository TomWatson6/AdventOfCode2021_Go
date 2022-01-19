
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func get_bit(lines []string, index int, mostCommon bool) rune {
	zero := 0
	one := 0

	for _, line := range lines {
		if []rune(line)[index] == '0' {
			zero += 1
		} else {
			one += 1
		}
	}

	if one >= zero {
		if mostCommon {
			return '1'
		} else {
			return '0'
		}
	} else {
		if mostCommon {
			return '0'
		} else {
			return '1'
		}
	}
}

func get_valid(lines []string, index int, char rune) []string {
	var ret []string

	for _, line := range lines {
		if []rune(line)[index] == char {
			ret = append(ret, line)
		}
	}

	return ret
}

func delete_empty(s []string) []string {
	var r []string

	for _, line := range s {
		if line != "" {
			r = append(r, line)
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

	var gamma []rune
	var epsilon []rune

	for x := 0; x < len(lines[0]); x++ {
		gamma = append(gamma, get_bit(lines, x, true))
		epsilon = append(epsilon, get_bit(lines, x, false))
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	e, _ := strconv.ParseInt(string(epsilon), 2, 64)

	fmt.Printf("Part 1: %d\n", g * e)

	og := lines
	cs := lines

	for x := 0; x < len(og[0]); x++ {
		mostCommon := get_bit(og, x, true)
		var og_ []string
		for _, line := range og {
			if []rune(line)[x] == mostCommon {
				og_ = append(og_, line)
			}
		}
		og = og_
		if len(og) == 1 {
			break
		}
	}

	for x := 0; x < len(cs[0]); x++ {
		leastCommon := get_bit(cs, x, false)
		var cs_ []string
		for _, line := range cs {
			if []rune(line)[x] == leastCommon {
				cs_ = append(cs_, line)
			}
		}
		cs = cs_
		if len(cs) == 1 {
			break
		}
	}

	og_value, _ := strconv.ParseInt(og[0], 2, 64)
	cs_value, _ := strconv.ParseInt(cs[0], 2, 64)

	fmt.Printf("Part 2: %d\n", og_value * cs_value)
}

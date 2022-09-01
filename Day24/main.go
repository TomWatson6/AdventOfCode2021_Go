package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var commands []command

var found []string

type computer struct {
	w, x, y, z *int
	mem        map[state]computer
}

func evaluate(values map[string]int, value string) int {
	if value == "" {
		return 0
	}

	if v, ok := values[value]; ok {
		return v
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return v
}

type state struct {
	depth, w, x, y, z int
}

type outcome struct {
	out bool
	rem string
}

var DP = make(map[state]outcome)

func search(trace string, depth, opIndex, wValue, xValue, yValue, zValue int) (bool, string) {
	s := state{
		depth: depth,
		w:     wValue,
		x:     xValue,
		y:     yValue,
		z:     zValue,
	}

	if outcome, ok := DP[s]; ok {
		return outcome.out, outcome.rem
	}

	if zValue > int(math.Pow10(7)) {
		DP[s] = outcome{false, ""}

		return false, ""
	}

	if opIndex >= len(commands) {
		DP[s] = outcome{zValue == 0, ""}

		if zValue == 0 {
			found = append(found, trace)
		}

		return zValue == 0, ""
	}

	values := map[string]int{
		"w": wValue,
		"x": xValue,
		"y": yValue,
		"z": zValue,
	}

	c := commands[opIndex]
	value := evaluate(values, c.right)

	switch c.code {
	case inp:
		for i := 1; i < 10; i++ {
			values[c.left] = i
			out, rem := search(trace+strconv.Itoa(i), depth+1, opIndex+1, values["w"], values["x"], values["y"], values["z"])

			if out {
				fmt.Printf("%d %d %d %d -> %s\n", wValue, xValue, yValue, zValue, trace)
				return out, strconv.Itoa(i) + rem
			}
		}

		return false, ""
	case add:
		values[c.left] += value
	case mul:
		values[c.left] *= value
	case div:
		val := math.Trunc(float64(values[c.left]) / float64(value))
		values[c.left] = int(val)
	case mod:
		values[c.left] %= value
	case eql:
		if values[c.left] == value {
			values[c.left] = 1
		} else {
			values[c.left] = 0
		}
	}

	out, rem := search(trace, depth+1, opIndex+1, values["w"], values["x"], values["y"], values["z"])
	DP[s] = outcome{out, rem}
	return out, rem
}

type opcode byte

const (
	inp opcode = iota
	add
	mul
	div
	mod
	eql
)

func parseOpCode(input string) opcode {
	switch input {
	case "inp":
		return inp
	case "add":
		return add
	case "mul":
		return mul
	case "div":
		return div
	case "mod":
		return mod
	case "eql":
		return eql
	default:
		panic("Unknown opcode: " + input)
	}
}

type command struct {
	code        opcode
	left, right string
}

func feed(cs []command, input int64) []command {
	str := []rune(strconv.FormatInt(input, 10))
	ptr := 0

	for i := range cs {
		if cs[i].code == inp {
			cs[i].right = string(str[ptr])
			ptr++
		}
	}

	return cs
}

func getInput(fileName string) []command {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fileString := string(bytes)
	lines := strings.Split(fileString, "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	var commands []command

	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		code := parseOpCode(splitLine[0])
		right := ""
		if len(splitLine) > 2 {
			right = splitLine[2]
		}

		commands = append(commands, command{code: code, left: splitLine[1], right: right})
	}

	return commands
}

func reduce(n int64) int64 {
	for {
		n--

		containsZero := false

		str := strconv.FormatInt(n, 10)
		for _, char := range str {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			if i == 0 {
				containsZero = true
			}
		}

		if !containsZero {
			return n
		}
	}
}

func main() {
	commands = getInput("input.txt")

	search("", 0, 0, 0, 0, 0, 0)

	fmt.Printf("Part 1: %v\n", found)
}

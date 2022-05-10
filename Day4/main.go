package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	Y, X int
}

type Card struct {
	Width   int
	Height  int
	Numbers [][]int
	Marked  map[Coord]bool
	HasWon  bool
}

type Winner struct {
	Card   Card
	Number int
}

func (c Card) GetSumUnmarked() int {
	sum := 0
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			if !c.Marked[Coord{y, x}] {
				sum += c.Numbers[y][x]
			}
		}
	}

	return sum
}

func (c *Card) Mark(num int) {
	for y, line := range c.Numbers {
		for x, value := range line {
			if value == num {
				c.Marked[Coord{y, x}] = true
				return
			}
		}
	}
}

func (c Card) Print() {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			coord := Coord{y, x}
			if c.Marked[coord] {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", ".")
			}
		}
		fmt.Printf("\n")
	}
}

func (c *Card) IsWinner() bool {
	//Go through horizontal lines
	for y := 0; y < c.Height; y++ {
		winner := true
		for x := 0; x < c.Width; x++ {
			if !c.Marked[Coord{y, x}] {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	//Go through vertical lines
	for x := 0; x < c.Width; x++ {
		winner := true
		for y := 0; y < c.Height; y++ {
			if !c.Marked[Coord{y, x}] {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	return false
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

func getNumbers(numberString string) []int {
	numberStrings := strings.Split(numberString, ",")
	var output []int

	for _, num := range numberStrings {
		number, _ := strconv.Atoi(num)
		output = append(output, number)
	}

	return output
}

func getCards(lines []string) []Card {
	var cards []Card
	var grid [][]int

	for _, line := range lines {
		if line == "" {
			cards = append(cards, Card{len(grid[0]), len(grid), grid, make(map[Coord]bool), false})
			grid = nil
			continue
		}
		var nums []int
		splitLine := strings.Split(line, " ")

		for _, s := range splitLine {
			if s != "" {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums = append(nums, num)
			}
		}
		grid = append(grid, nums)
	}

	return cards
}

func main() {
	lines := getInput("input.txt")

	numbers := getNumbers(lines[0])
	cards := getCards(lines[2:])

	var winners []Winner

	for _, number := range numbers {
		for i := 0; i < len(cards); i++ {
			if !cards[i].HasWon {
				cards[i].Mark(number)
			}
		}

		for i := 0; i < len(cards); i++ {
			if cards[i].IsWinner() && !cards[i].HasWon {
				cards[i].HasWon = true
				winners = append(winners, Winner{cards[i], number})
			}
		}
	}

	fmt.Printf("Part 1: %d\n", winners[0].Number*winners[0].Card.GetSumUnmarked())
	fmt.Printf("Part 2: %d\n", winners[len(winners)-1].Number*winners[len(winners)-1].Card.GetSumUnmarked())
}

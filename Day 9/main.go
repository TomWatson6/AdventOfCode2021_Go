package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Coord2 struct {
	R, C int
}

func getInput(fileName string) [][]int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}

	var output [][]int

	for _, line := range lines {
		var numbers []int

		lineChars := []rune(line)

		for _, char := range lineChars {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, num)
		}

		output = append(output, numbers)
	}

	return output
}

func getPits(grid [][]int) []int {
	var collection []Coord2
	var pits []int

	//Locate the centre of the pits
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if r == len(grid)-1 || grid[r][c] < grid[r+1][c] {
				if r == 0 || grid[r][c] < grid[r-1][c] {
					if c == len(grid[0])-1 || grid[r][c] < grid[r][c+1] {
						if c == 0 || grid[r][c] < grid[r][c-1] {
							collection = append(collection, Coord2{r, c})
						}
					}
				}
			}
		}
	}

	//For each pit, find surrounding parts that need to be checked, and add them to a finitely expanding list.
	//The expansion will stop once the pits edges have been found (9 has been found around the edges)
	for _, c := range collection {
		var toCheck []Coord2
		toCheck = append(toCheck, c)
		curr := 0

		for {
			if curr == len(toCheck) {
				break
			}
			r := toCheck[curr].R
			c := toCheck[curr].C

			if r != len(grid)-1 {
				if grid[r][c] < grid[r+1][c] {
					if grid[r+1][c] != 9 {
						toCheck = append(toCheck, Coord2{r + 1, c})
					}
				}
			}
			if r != 0 {
				if grid[r][c] < grid[r-1][c] {
					if grid[r-1][c] != 9 {
						toCheck = append(toCheck, Coord2{r - 1, c})
					}
				}
			}
			if c != len(grid[0])-1 {
				if grid[r][c] < grid[r][c+1] {
					if grid[r][c+1] != 9 {
						toCheck = append(toCheck, Coord2{r, c + 1})
					}
				}
			}
			if c != 0 {
				if grid[r][c] < grid[r][c-1] {
					if grid[r][c-1] != 9 {
						toCheck = append(toCheck, Coord2{r, c - 1})
					}
				}
			}
			curr += 1
		}
		strippedCheck := make(map[Coord2]int)
		for _, t := range toCheck {
			strippedCheck[t] += 1
		}
		count := len(strippedCheck)
		pits = append(pits, count)
	}

	return pits
}

func main() {
	grid := getInput("input.txt")
	total := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if r == len(grid)-1 || grid[r][c] < grid[r+1][c] {
				if r == 0 || grid[r][c] < grid[r-1][c] {
					if c == len(grid[0])-1 || grid[r][c] < grid[r][c+1] {
						if c == 0 || grid[r][c] < grid[r][c-1] {
							total += 1 + grid[r][c]
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total)

	pits := getPits(grid)
	sort.Ints(pits)

	largest := 1

	for i := len(pits) - 1; i >= len(pits)-3; i-- {
		largest *= pits[i]
	}

	fmt.Printf("Part 2: %d\n", largest)
}


package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(nums []string, size int) int {
	output := 0

	for i := 0; i < len(nums) - size; i++ {
		left, _ := strconv.Atoi(nums[i])
		right, _ := strconv.Atoi(nums[i + size])

		if left < right {
			output += 1
		}
	}

	return output
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)
	components := strings.Split(fileString, "\n")

	fmt.Printf("Part 1: %d\n", solve(components, 1))
	fmt.Printf("Part 2: %d\n", solve(components, 3))
}

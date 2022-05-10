package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Node string
type Nodes []Node

func (n Node) IsLower() bool {
	for _, c := range n {
		if unicode.IsLetter(c) && !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func (n Node) IsUpper() bool {
	for _, c := range n {
		if unicode.IsLetter(c) && !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

func (ns Nodes) Contains(n Node) bool {
	for _, n1 := range ns {
		if Node(n1) == n {
			return true
		}
	}
	return false
}

func findPaths(nodes map[Node]Nodes, currentNode Node, currentPath Nodes, extraSmallVisited bool) int {
	if currentNode == "end" {
		return 1
	}

	var terminalNodes Nodes
	terminalNodes = append(terminalNodes, "start")
	terminalNodes = append(terminalNodes, "end")

	numPaths := 0
	next := nodes[currentNode]
	var toExplore Nodes
	for _, n := range next {
		if n.IsUpper() || !currentPath.Contains(n) {
			toExplore = append(toExplore, n)
		} else if !extraSmallVisited && !terminalNodes.Contains(n) {
			toExplore = append(toExplore, n)
		}
	}
	if len(toExplore) == 0 {
		return 0
	} else {
		previousNodes := currentPath
		previousNodes = append(previousNodes, currentNode)
		for _, n := range toExplore {
			if currentPath.Contains(n) && n.IsLower() {
				numPaths += findPaths(nodes, n, previousNodes, true)
			} else {
				numPaths += findPaths(nodes, n, previousNodes, extraSmallVisited)
			}
		}

		return numPaths
	}
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

	var output []string

	for _, line := range lines {
		if line != "" {
			output = append(output, line)
		}
	}

	return output
}

func parseInput(lines []string) map[Node]Nodes {
	nodes := make(map[Node]Nodes)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		from := Node(parts[0])
		to := Node(parts[1])

		nodes[from] = append(nodes[from], to)
		nodes[to] = append(nodes[to], from)
	}

	return nodes
}

func main() {
	lines := getInput("input.txt")
	nodes := parseInput(lines)

	paths := findPaths(nodes, "start", Nodes{}, true)
	fmt.Printf("Part 1: %d\n", paths)

	paths = findPaths(nodes, "start", Nodes{}, false)
	fmt.Printf("Part 2: %d\n", paths)
}

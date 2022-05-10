package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calculate(values []int, outerId int) int {
	switch outerId {
	case 0:
		sum := 0
		for _, value := range values {
			sum += value
		}
		return sum
	case 1:
		product := 1
		for _, value := range values {
			product *= value
		}
		return product
	case 2:
		min := 999999
		for _, value := range values {
			if value < min {
				min = value
			}
		}
		return min
	case 3:
		max := 0
		for _, value := range values {
			if value > max {
				max = value
			}
		}
		return max
	case 5:
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	default:
		return -1
	}
}

func decode(binary []rune, outerId int, terminationLength int) (int, int, []rune) {
	versionTotal := 0
	var values []int
	numToCollect := terminationLength
	packetsRemaining := terminationLength

	for {
		if len(binary) <= 7 {
			break
		}
		version, err := strconv.ParseInt(string(binary[:3]), 2, 64)
		if err != nil {
			panic(err)
		}
		binary = binary[3:]
		versionTotal += int(version)

		typeId, _ := strconv.ParseInt(string(binary[:3]), 2, 64)
		binary = binary[3:]

		if typeId == 4 {
			value := ""
			for {
				segment := binary[:5]
				binary = binary[5:]
				for _, char := range segment[1:] {
					value += string(char)
				}
				if string(segment[0]) == "0" {
					break
				}
			}
			val, _ := strconv.ParseInt(value, 2, 64)
			values = append(values, int(val))
		} else {
			lengthTypeId, _ := strconv.Atoi(string(binary[0]))
			binary = binary[1:]
			if lengthTypeId == 0 {
				length, _ := strconv.ParseInt(string(binary[:15]), 2, 64)
				binary = binary[15:]
				versionValue, outcome, _ := decode(binary[:length], int(typeId), 0)
				versionTotal += versionValue
				values = append(values, outcome)
				binary = binary[length:]
			} else {
				amount, _ := strconv.ParseInt(string(binary[:11]), 2, 64)
				binary = binary[11:]
				versionValue, outcome, newBin := decode(binary, int(typeId), int(amount))
				versionTotal += versionValue
				values = append(values, outcome)
				binary = newBin
			}
		}

		if packetsRemaining > 0 {
			packetsRemaining -= 1
		}

		if packetsRemaining == 0 && numToCollect > 0 {
			break
		}
	}

	return versionTotal, calculate(values, outerId), binary
}

func getInput(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(bytes)
	lines := strings.Split(fileString, "\n")

	var output []string

	for _, line := range lines {
		if line != "" {
			line = strings.Trim(line, "\r")
			output = append(output, line)
		}
	}

	return output[0]
}

func main() {
	bin := make(map[rune]string)
	bin['0'] = "0000"
	bin['1'] = "0001"
	bin['2'] = "0010"
	bin['3'] = "0011"
	bin['4'] = "0100"
	bin['5'] = "0101"
	bin['6'] = "0110"
	bin['7'] = "0111"
	bin['8'] = "1000"
	bin['9'] = "1001"
	bin['A'] = "1010"
	bin['B'] = "1011"
	bin['C'] = "1100"
	bin['D'] = "1101"
	bin['E'] = "1110"
	bin['F'] = "1111"

	hex := getInput("input.txt")
	binary := ""

	for _, char := range []rune(hex) {
		binary += bin[char]
	}

	versionTotal, outcome, _ := decode([]rune(binary), 0, 0)

	fmt.Printf("Part 1: %d\n", versionTotal)
	fmt.Printf("Part 2: %d\n", outcome)

}

/*
Advent of Code 2020 day 08
Seven Segment Search
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

// Count 1, 4, 7, 8 in output values
func SolvePart1(input []string) int {
	_, outputs := ParseInput(input)
	digits := GetDigits(&outputs)

	// Add easy numbers
	sum := 0
	for _, set := range digits {
		for _, v := range set {
			if v > -1 {
				sum++
			}
		}
	}

	return sum
}

// ???
func SolvePart2(input []string) int {
	return -1
}

func ParseInput(input []string) (inputs, outputs [][]string) {
	for _, line := range input {
		line_split := strings.Split(line, " | ")
		inputs = append(inputs, strings.Fields(line_split[0]))
		outputs = append(outputs, strings.Fields(line_split[1]))
	}

	return inputs, outputs
}

func GetDigits(input *[][]string) [][]int {
	digits := make([][]int, len(*input))

	for i, line := range *input {
		digits[i] = make([]int, len((*input)[i]))
		for j, set := range line {
			switch len(set) {
			case 2:
				digits[i][j] = 1
			case 3:
				digits[i][j] = 7
			case 4:
				digits[i][j] = 4
			case 7:
				digits[i][j] = 8
			default:
				digits[i][j] = -1
			}
		}
	}

	return digits
}

func main() {
	input := shared.ReadInputLines("day08/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

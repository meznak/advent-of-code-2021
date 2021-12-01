/*
Advent of Code 2021 day 01
Sonar Sweep
*/

package main

import (
	"aoc2021/shared"
	"fmt"
)

// Count depth increases
func SolvePart1(input []int) int {
	var prev_depth, increases int

	for i, depth := range input {
		// First measurement can't be deeper
		if i == 0 {
			prev_depth = depth
			continue
		}

		if depth > prev_depth {
			increases++
		}

		prev_depth = depth
	}

	return increases
}

// Count rolling sum increases
func SolvePart2(input []int) int {
	sums := SumWindows(input)
	return SolvePart1(sums)
}

// Sum a three-measurement rolling window
func SumWindows(input []int) (sums []int) {
	for i := range input {
		// stop when the last value has been read
		if i < len(input)-2 {
			sums = append(sums, input[i]+input[i+1]+input[i+2])
		}
	}

	return
}

func main() {
	input := shared.ReadInputInts("day01/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

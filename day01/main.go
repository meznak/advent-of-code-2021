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

// ???
func SolvePart2(input []int) int {

	return -1
}

func main() {
	input := shared.ParseInputInts(shared.ReadInputLines("./input"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

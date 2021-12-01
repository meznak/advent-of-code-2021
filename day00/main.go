/*
Advent of Code 2020 day 01
Report Repair
*/

package main

import (
	"aoc2021/shared"
	"fmt"
)

// Find pairs that add to 2020
func SolvePart1(input []int) int {
	for i, a := range input {
		for _, b := range input[i:] {
			if a+b == 2020 {
				return a + b
			}
		}
	}

	return -1
}

// Find triples that add to 2020
func SolvePart2(input []int) int {
	for i, a := range input {
		for j, b := range input[i:] {
			for _, c := range input[j:] {
				if a+b+c == 2020 {
					return a + b
				}
			}
		}
	}

	return -1
}

func main() {
	input := shared.ParseInputInts(shared.ReadInputLines("./input"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

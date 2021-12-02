/*
Advent of Code 2020 day 02
Dive!
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

// Calculate horizontal position & depth
func SolvePart1(input []string) int {
	var x, z int

	for _, line := range input {
		instruction := strings.Split(line, " ")
		distance, err := strconv.Atoi(instruction[1])

		if err != nil {
			panic(err)
		}

		switch instruction[0] {
		case "forward":
			x += distance
		case "down":
			z += distance
		case "up":
			z -= distance
		}
	}

	// Return horizontal * depth
	return x * z
}

// Find triples that add to 2020
func SolvePart2(input []string) int {

	return -1
}

func main() {
	input := shared.ReadInputLines("day02/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

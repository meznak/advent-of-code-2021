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

type instruction struct {
	direction byte // first character
	distance  int
}

// Calculate horizontal position & depth
func SolvePart1(input []string) int {
	var x, z int

	instructions := ParseInstructions(&input)

	for _, inst := range instructions {

		switch inst.direction {
		case 'f': // forward
			x += inst.distance
		case 'd': // down
			z += inst.distance
		case 'u': // up
			z -= inst.distance
		}
	}

	// Return horizontal * depth
	return x * z
}

// Find triples that add to 2020
func SolvePart2(input []string) int {
	var x, z, aim int

	instructions := ParseInstructions(&input)

	for _, inst := range instructions {

		switch inst.direction {
		case 'f':
			x += inst.distance
			z += inst.distance * aim
		case 'd':
			aim += inst.distance
		case 'u':
			aim -= inst.distance
		}
	}

	// Return horizontal * depth
	return x * z
}

func ParseInstructions(input *[]string) []instruction {
	out := make([]instruction, len(*input))

	for i, line := range *input {
		direction := line[0]
		distance, err := strconv.Atoi(line[strings.LastIndex(line, " ")+1:])

		if err != nil {
			panic(err)
		}

		out[i] = instruction{direction, distance}

	}

	return out
}

func main() {
	input := shared.ReadInputLines("day02/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

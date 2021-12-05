/*
Advent of Code 2020 day 03
Binary Diagnostic
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

// Count line intersections
func SolvePart1(input []string) int {
	const grid_size int = 9
	// Create grid
	grid := make([][grid_size]int, grid_size)

	// Parse lines
	lines := ParseLines(input)
	// Apply lines
	for _, line := range lines {

		switch {
		case line[2] > line[0]: // right
			for x := line[0]; x < line[2]; x++ {
				grid[x][line[1]]++
			}
		case line[3] > line[1]: // down
			for x := line[1]; x < line[3]; x++ {
				grid[x][line[0]]++
			}
		case line[2] > line[0]: // left
			for x := line[0]; x < line[2]; x++ {
				grid[x][line[1]]++
			}
		case line[3] > line[1]: // up
			for x := line[0]; x < line[2]; x++ {
				grid[x][line[0]]++
			}
		}

	}

	// Count intersections

	return -1
}

// ???
func SolvePart2(input []string) int {

	return -1
}

func ParseLines(input []string) (out [][]int) {
	// 0,9 -> 5,9
	for _, line := range input {
		move := make([]int, 0)
		move_s := strings.Split(line, " -> ")
		start_s := strings.Split(move_s[0], ",")
		end_s := strings.Split(move_s[1], ",")

		for _, v := range start_s {
			val, _ := strconv.Atoi(v)
			move = append(move, val)
		}

		for _, v := range end_s {
			val, _ := strconv.Atoi(v)
			move = append(move, val)
		}
		out = append(out, move)
	}

	return
}

func main() {
	input := shared.ReadInputLines("day04/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

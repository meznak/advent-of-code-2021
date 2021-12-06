/*
Advent of Code 2020 day 06
Lanternfish
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

// Count exponentially spawning fish
func SolvePart1(input []string) int {
	// Parse list
	var fishes []int
	input_s := strings.Split(input[0], ",")
	for _, v := range input_s {
		fish, _ := strconv.Atoi(v)
		fishes = append(fishes, fish)
	}

	// Update list
	for day := 0; day < 80; day++ {
		ProcessDay(&fishes)
	}

	return len(fishes)
}

// Same for 256 days
func SolvePart2(input []string) int {
	return -1
}

func ProcessDay(fishes *[]int) {
	new_fish := 0

	// Decrease timers
	for i, v := range *fishes {
		if v == 0 {
			new_fish++
			(*fishes)[i] = 6
		} else {
			(*fishes)[i]--
		}
	}

	// Add new fish
	for new_fish > 0 {
		(*fishes) = append((*fishes), 8)
		new_fish--
	}
}

func main() {
	input := shared.ReadInputLines("day06/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

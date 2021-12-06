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
	grouped_fish := GroupFish(&fishes)

	// Update list
	for day := 0; day < 80; day++ {
		ProcessDay(&grouped_fish)
	}

	// Count fish
	count := 0
	for _, v := range grouped_fish {
		count += v
	}

	return count
}

// Same for 256 days
func SolvePart2(input []string) int {
	// Parse list
	var fishes []int
	input_s := strings.Split(input[0], ",")
	for _, v := range input_s {
		fish, _ := strconv.Atoi(v)
		fishes = append(fishes, fish)
	}
	grouped_fish := GroupFish(&fishes)

	// Update list
	for day := 0; day < 256; day++ {
		ProcessDay(&grouped_fish)
	}

	// Count fish
	count := 0
	for _, v := range grouped_fish {
		count += v
	}

	return count
}

func GroupFish(fishes *[]int) []int {
	// Count fish by days remaining to simplify the slice

	groups := make([]int, 9)

	for _, v := range *fishes {
		groups[v]++
	}

	return groups
}

func ProcessDay(fishes *[]int) {
	// Decrease timers
	new_fish := (*fishes)[0]

	for i, v := range *fishes {
		if i > 0 {
			(*fishes)[i-1] = v
		}
	}

	// Add new parents
	(*fishes)[6] += new_fish

	// Add new fish
	(*fishes)[8] = new_fish
}

func main() {
	input := shared.ReadInputLines("day06/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

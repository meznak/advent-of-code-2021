/*
Advent of Code 2020 day 07
The Treachery of Whales
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"math"
)

// Calculate fuel to align crabs
func SolvePart1(input []string) []int {
	crabs := shared.CSVToInt(input[0])
	fuel := CalculateFuel(&crabs)

	total_fuel := SumFuel(&fuel)

	best_position := MinFuel(&total_fuel)

	return []int{best_position, total_fuel[best_position]}
}

// ???
func SolvePart2(input []string) int {

	return -1
}

// Calculate fuel from each crab to each position
func CalculateFuel(crabs *[]int) [][]int {
	num_crabs := len(*crabs)
	fuel := make([][]int, num_crabs)

	for i, crab_a := range *crabs {
		fuel[i] = make([]int, num_crabs)
		for j, crab_b := range *crabs {
			fuel[i][j] = AbsInt(crab_a - crab_b)
		}
	}

	return fuel
}

// Find total fuel to each position
func SumFuel(fuel *[][]int) []int {
	totals := make([]int, len(*fuel))
	for j := range *fuel {
		// position
		for i := range *fuel {
			// crab
			totals[i] += (*fuel)[i][j]
		}
	}

	return totals
}

func AbsInt(number int) int {
	if number < 0 {
		return -1 * number
	}

	return number
}

func MinFuel(fuel *[]int) int {
	lowest_i := 0
	lowest_v := int(math.MaxInt64)
	for i, v := range *fuel {
		if v < lowest_v {
			lowest_i = i
			lowest_v = v
		}
	}
	return lowest_i
}

func main() {
	input := shared.ReadInputLines("day07/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

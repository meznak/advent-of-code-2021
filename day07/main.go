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
	fuel := CalculateFuel_Part1(&crabs)

	total_fuel := SumFuel(&fuel)

	best_position := MinFuel(&total_fuel)

	return []int{best_position, total_fuel[best_position]}
}

// Same with different fuel calculation
func SolvePart2(input []string) []int {
	crabs := shared.CSVToInt(input[0])
	fuel := CalculateFuel_Part2(&crabs)

	total_fuel := SumFuel(&fuel)

	best_position := MinFuel(&total_fuel)

	return []int{best_position, total_fuel[best_position]}
}

// Calculate fuel from each crab to each position
func CalculateFuel_Part1(crabs *[]int) [][]int {
	num_positions := len(*crabs)
	fuel := make([][]int, num_positions)

	for i, crab_a := range *crabs {
		fuel[i] = make([]int, num_positions)
		for j, crab_b := range *crabs {
			fuel[i][j] = AbsInt(crab_a - crab_b)
		}
	}

	return fuel
}

// Calculate fuel from each crab to each position
// Each step farther costs 1 more fuel per step
func CalculateFuel_Part2(crabs *[]int) [][]int {
	num_positions := MaxInt(crabs) + 1
	fuel := make([][]int, len(*crabs))
	var prev_fuel, delta, new_fuel int

	for i, crab := range *crabs {
		fuel[i] = make([]int, num_positions)
		prev_fuel, delta, new_fuel = 0, 0, 0
		for j := crab - 1; j >= 0; j-- {
			delta++
			new_fuel = delta + prev_fuel
			fuel[i][j] = new_fuel
			prev_fuel = new_fuel
		}
		prev_fuel, delta, new_fuel = 0, 0, 0
		for j := crab + 1; j < num_positions; j++ {
			delta++
			new_fuel = delta + prev_fuel
			fuel[i][j] = new_fuel
			prev_fuel = new_fuel
		}
	}

	return fuel
}

// Find total fuel to each position
func SumFuel(fuel *[][]int) []int {
	totals := make([]int, len((*fuel)[0]))
	for j := range (*fuel)[0] {
		// position
		for i := range *fuel {
			// crab
			totals[j] += (*fuel)[i][j]
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

func MaxInt(numbers *[]int) int {
	max := 0
	for _, v := range *numbers {
		if v > max {
			max = v
		}
	}

	return max
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

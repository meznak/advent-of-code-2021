/*
Advent of Code 2020 day 07
The Treachery of Whales
*/

package main

import (
	"aoc2021/shared"
	"fmt"
)

// Calculate fuel to align crabs
func SolvePart1(input []string) []int {
	crabs := shared.CSVToInt(input[0])
	fuel := CalculateFuel_Part1(&crabs)

	total_fuel := SumFuel(&fuel)

	best_position, best_fuel := shared.MinInt(&total_fuel)

	return []int{best_position, best_fuel}
}

// Same with different fuel calculation
func SolvePart2(input []string) []int {
	crabs := shared.CSVToInt(input[0])
	fuel := CalculateFuel_Part2(&crabs)

	total_fuel := SumFuel(&fuel)

	best_position, best_fuel := shared.MinInt(&total_fuel)

	return []int{best_position, best_fuel}
}

// Calculate fuel from each crab to each position
func CalculateFuel_Part1(crabs *[]int) [][]int {
	num_positions := len(*crabs)
	fuel := make([][]int, num_positions)

	for i, crab_a := range *crabs {
		fuel[i] = make([]int, num_positions)
		for j, crab_b := range *crabs {
			fuel[i][j] = shared.AbsInt(crab_a - crab_b)
		}
	}

	return fuel
}

// Calculate fuel from each crab to each position
// Each step farther costs 1 more fuel per step
func CalculateFuel_Part2(crabs *[]int) [][]int {
	_, num_positions := shared.MaxInt(crabs)
	num_positions++
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
	ch := make(chan int)
	totals := make([]int, len((*fuel)[0]))
	for col := range (*fuel)[0] {
		go shared.SumVerticalInt(fuel, col, ch)
		totals[col] = <-ch
	}

	return totals
}

func main() {
	input := shared.ReadInputLines("day07/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

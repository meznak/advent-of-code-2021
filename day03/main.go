/*
Advent of Code 2020 day 03
Binary Diagnostic
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
)

// Convert binary to "gamma" and "epsilon" numbers
func SolvePart1(input []string) int {
	length := len(input[0])
	var g, e string

	for i_bit := 0; i_bit < length; i_bit++ {
		var ones, zeros int

		for _, byte := range input {
			if byte[i_bit] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		if ones > zeros {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}

	g_i, _ := strconv.ParseUint(g, 2, 64)
	e_i, _ := strconv.ParseUint(e, 2, 64)

	return int(g_i * e_i)
}

//
func SolvePart2(input []string) int {

	return -1
}

func main() {
	input := shared.ReadInputLines("day03/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

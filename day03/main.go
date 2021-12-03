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
	mcb := MostCommonBits(input)
	lcb := FlipBits(mcb)

	g_i, _ := strconv.ParseUint(mcb, 2, 64)
	e_i, _ := strconv.ParseUint(lcb, 2, 64)

	return int(g_i * e_i)
}

//
func SolvePart2(input []string) int {
	var mcb byte
	oxy := make([]string, len(input))
	co2 := make([]string, len(input))

	copy(oxy, input)
	for i := range input[0] {
		mcb = MostCommonBit(oxy, i)

		for j := len(oxy) - 1; j >= 0; j-- {
			if oxy[j][i] != mcb {
				oxy[j] = oxy[len(oxy)-1]
				oxy = oxy[:len(oxy)-1]
			}
		}

		if len(oxy) == 1 {
			break
		}
	}

	copy(co2, input)
	for i := range input[0] {
		mcb := MostCommonBit(co2, i)

		for j := len(co2) - 1; j >= 0; j-- {
			if co2[j][i] == mcb {
				co2[j] = co2[len(co2)-1]
				co2 = co2[:len(co2)-1]
			}
		}

		if len(co2) == 1 {
			break
		}
	}

	oxy_i, _ := strconv.ParseUint(oxy[0], 2, 64)
	co2_i, _ := strconv.ParseUint(co2[0], 2, 64)

	return int(oxy_i * co2_i)
}

func MostCommonBit(input []string, i_bit int) byte {
	count := 0

	for _, line := range input {
		if line[i_bit] == '1' {
			count++
		} else {
			count--
		}
	}

	if count >= 0 {
		return '1'
	} else {
		return '0'
	}
}

func MostCommonBits(input []string) (out string) {
	length := len(input[0])

	for i_bit := 0; i_bit < length; i_bit++ {
		out += string(MostCommonBit(input, i_bit))
	}

	return
}

func FlipBits(input string) (out string) {
	for _, v := range input {
		if v == '0' {
			out += "1"
		} else {
			out += "0"
		}
	}

	return
}

func main() {
	input := shared.ReadInputLines("day03/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

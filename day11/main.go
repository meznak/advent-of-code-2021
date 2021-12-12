/*
Advent of Code 2020 day 11
Dumbo Octopus
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
)

var (
	flash_count int
)

// Count octopus flashes
func SolvePart1(input []string) int {
	octos := ParseInput(&input)
	rows := len(octos)
	cols := len(octos[0])
	var changed bool

	for step := 0; step < 100; step++ {
		// Grid to track flashes
		flashed := make([][]bool, rows)

		for y := range octos {
			flashed[y] = make([]bool, cols)
		}

		for y, row := range octos {
			for x := range row {
				if !flashed[y][x] {
					octos[y][x]++
				}
			}
		}

		changed = true
		for changed {
			changed = false
			for y, row := range octos {
				for x, octo := range row {
					if octo > 9 {
						changed = true
						Flash(&octos, &flashed, y, x)
					}
				}
			}
		}
	}

	return flash_count
}

// ???
func SolvePart2(input []string) int {
	return -1
}

func ParseInput(input *[]string) [][]int {
	ints := make([][]int, len(*input))
	line_length := len((*input)[0])

	for i, line := range *input {
		temp := make([]int, line_length)
		for j, v := range line {
			temp[j], _ = strconv.Atoi(string(v))
		}

		ints[i] = temp
	}

	return ints
}

func Update(octos *[][]int, flashed *[][]bool, y, x int) {
	if (*flashed)[x][y] {
		// Can't flash twice in one step
		return
	}

	(*octos)[y][x]++

	if (*octos)[y][x] > 9 {
		Flash(octos, flashed, y, x)
	}
}

func Flash(octos *[][]int, flashed *[][]bool, y, x int) {
	flash_count++
	(*flashed)[y][x] = true
	(*octos)[y][x] = 0

	// Increment neighbors
	for yn := y - 1; yn <= y+1; yn++ {
		for xn := x - 1; xn <= x+1; xn++ {
			// check bounds and ignore self
			if yn < 0 || xn < 0 || (yn == y && xn == x) || yn >= len(*octos) || xn >= len((*octos)[0]) {
				continue
			}

			//Update(octos, flashed, yn, xn)
			if !(*flashed)[yn][xn] {
				(*octos)[yn][xn]++
			}
		}
	}
}

func main() {
	input := shared.ReadInputLines("day11/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

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

type Vector struct {
	x int
	y int
}

// Count line intersections
func SolvePart1(input []string) int {
	points := make(map[Vector]int, 0)

	// Parse lines
	lines := ParseLines(input)

	// Apply lines
	ApplyLines(&points, &lines)

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
		if len(line) < 10 {
			continue
		}

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

func ApplyLines(points *map[Vector]int, lines *[][]int) {
	for _, line := range *lines {
		// x1 y1 x2 y2
		//  0  1  2  3

		switch {
		case line[2] > line[0]: // right
			for x := line[0]; x <= line[2]; x++ {
				point := Vector{x, line[1]}
				AddPoint(points, point)
			}
		case line[3] > line[1]: // down
			for y := line[1]; y <= line[3]; y++ {
				point := Vector{line[0], y}
				AddPoint(points, point)
			}
		case line[2] < line[0]: // left
			for x := line[0]; x >= line[2]; x-- {
				point := Vector{x, line[1]}
				AddPoint(points, point)
			}
		case line[3] < line[1]: // up
			for y := line[1]; y >= line[3]; y-- {
				point := Vector{line[0], y}
				AddPoint(points, point)
			}
		}
	}
}

func AddPoint(points *map[Vector]int, point Vector) {
	if _, exists := (*points)[point]; !exists {
		(*points)[point] = 0
	}

	(*points)[point]++
}

func main() {
	input := shared.ReadInputLines("day05/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

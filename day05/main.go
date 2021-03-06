/*
Advent of Code 2020 day 05
Hydrothermal Venture
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
	ApplyLines(&points, &lines, true)

	// Count intersections
	count := CountIntersections(&points)

	return count
}

// Same, considering diagnoal lines
func SolvePart2(input []string) int {
	points := make(map[Vector]int, 0)

	// Parse lines
	lines := ParseLines(input)

	// Apply lines
	ApplyLines(&points, &lines, false)

	// Count intersections
	count := CountIntersections(&points)

	return count
}

func ParseLines(input []string) (out [][]int) {
	// 0,9 -> 5,9
	// TODO: Refactor to use vectors
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

func ApplyLines(points *map[Vector]int, lines *[][]int, ignore_diag bool) {
	for _, line := range *lines {
		// x1 y1 x2 y2
		//  0  1  2  3

		if ignore_diag && line[0] != line[2] && line[1] != line[3] {
			continue
		}

		x_diff := line[2] - line[0]
		y_diff := line[3] - line[1]

		var x_d, y_d int

		switch {
		case x_diff < 0:
			x_d = -1
		case x_diff == 0:
			x_d = 0
		case x_diff > 0:
			x_d = 1
		}

		switch {
		case y_diff < 0:
			y_d = -1
		case y_diff == 0:
			y_d = 0
		case y_diff > 0:
			y_d = 1
		}

		x, y := line[0], line[1]
		for {
			AddPoint(points, Vector{x, y})
			if x == line[2] && y == line[3] {
				break
			}
			x += x_d
			y += y_d
		}
	}
}

func AddPoint(points *map[Vector]int, point Vector) {
	if _, exists := (*points)[point]; !exists {
		(*points)[point] = 0
	}

	(*points)[point]++
}

func CountIntersections(points *map[Vector]int) (count int) {
	for _, point_count := range *points {
		if point_count > 1 {
			count++
		}
	}

	return
}

func main() {
	input := shared.ReadInputLines("day05/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

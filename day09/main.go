/*
Advent of Code 2020 day 09
Smoke Basin
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
	"strconv"
)

// Find local minimum depths
func SolvePart1(input []string) int {
	heights := ParseInput(&input)
	minima := FindLowPoints(&heights)
	risk := CalcRisk(&heights, &minima)

	return risk
}

// Find basins around local minima
func SolvePart2(input []string) int {
	heights := ParseInput(&input)
	basins := FindBasins(&heights)

	sort.Slice(basins, func(i, j int) bool { return basins[j] < basins[i] })
	product := 1
	for _, v := range basins[:3] {
		product *= v
	}

	return product
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

// Check directions, respecting boundaries
func FindLowPoints(heights *[][]int) [][]bool {
	neighbors := make([]int, 4)
	minima := make([][]bool, len(*heights))
	row_length := len((*heights)[0])

	for y, row := range *heights {
		row_minima := make([]bool, row_length)

		for x, val := range row {
			neighbors = []int{10, 10, 10, 10}
			// up
			if y > 0 {
				neighbors[0] = (*heights)[y-1][x]
			}

			// right
			if x < len(row)-1 {
				neighbors[1] = (*heights)[y][x+1]
			}

			// down
			if y < len(*heights)-1 {
				neighbors[2] = (*heights)[y+1][x]
			}

			// left
			if x > 0 {
				neighbors[3] = (*heights)[y][x-1]
			}

			_, neighbor_min := shared.MinInt(&neighbors)

			if neighbor_min <= val {
				row_minima[x] = false
			} else {
				row_minima[x] = true
			}
		}

		minima[y] = row_minima
	}

	return minima
}

func CalcRisk(heights *[][]int, minima *[][]bool) (risk int) {
	for y, row := range *minima {
		for x, val := range row {
			if val {
				risk += (*heights)[y][x] + 1
			}
		}
	}

	return
}

func FindBasins(heights *[][]int) (basins []int) {
	// build memo to track whether a point has been checked
	visited := make([][]bool, len(*heights))
	for y := range visited {
		visited[y] = make([]bool, len((*heights)[0]))
	}

	for y, row := range *heights {
		for x := range row {
			if !visited[y][x] {
				basin_size := WalkBasin(heights, &visited, y, x)
				if basin_size > 0 {
					basins = append(basins, basin_size)
				}
			}
		}
	}

	return
}

func WalkBasin(heights *[][]int, visited *[][]bool, y, x int) (size int) {
	was_visited := (*visited)[y][x]
	(*visited)[y][x] = true

	if (*heights)[y][x] == 9 || was_visited {
		return 0
	}

	size += 1

	// up
	if y > 0 {
		size += WalkBasin(heights, visited, y-1, x)
	}

	// right
	if x < len((*heights)[0])-1 {
		size += WalkBasin(heights, visited, y, x+1)
	}

	// down
	if y < len(*heights)-1 {

		size += WalkBasin(heights, visited, y+1, x)
	}

	// left
	if x > 0 {

		size += WalkBasin(heights, visited, y, x-1)
	}

	return
}

func main() {
	input := shared.ReadInputLines("day09/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

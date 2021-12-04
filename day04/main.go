/*
Advent of Code 2020 day 03
Binary Diagnostic
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

// Convert binary to "gamma" and "epsilon" numbers
func SolvePart1(input []string) int {

	numbers := ParseNumbers(input[0])
	boards := ParseBoards(input[2:])

	board_state := make([][5][5]bool, len(boards))
	for i := range board_state {
		for j := 0; j < 5; j++ {
			board_state[i][j] = [5]bool{false, false, false, false, false}
		}
	}

	for _, number := range numbers {
		CheckBoards(number, &boards, &board_state)
	}

	return -1
}

//
func SolvePart2(input []string) int {
	return -1
}

func ParseNumbers(input string) (output []string) {
	return strings.Split(input, ",")
}

func ParseBoards(input []string) (output [][][]string) {
	board := make([][]string, 0)

	for _, line := range input {
		if line == "" {
			continue
		}

		board = append(board, strings.Fields(line))

		if len(board) == 5 {
			output = append(output, board)
			board = make([][]string, 0)
		}
	}

	return
}

func CheckBoards(number string, boards *[][][]string, board_state *[][5][5]bool) {
	for i, board := range *boards {
		for j := range board {
			for k := range board[0] {
				if board[j][k] == number {
					(*board_state)[i][j][k] = true
				}
			}
		}
	}
}

func main() {
	input := shared.ReadInputLines("day03/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

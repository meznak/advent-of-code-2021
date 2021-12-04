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

// Convert binary to "gamma" and "epsilon" numbers
func SolvePart1(input []string) int {
	var winner int

	numbers := ParseNumbers(input[0])
	boards := ParseBoards(input[2:])

	board_state := make([][5][5]bool, len(boards))
	/* 	for i := range board_state {
		for j := 0; j < 5; j++ {
			board_state[i][j] = [5]bool{false, false, false, false, false}
		}
	} */

	for _, number := range numbers {
		MarkBoards(number, &boards, &board_state)
		winner = CheckBoards(&board_state)
		if winner > -1 {
			break
		}
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

func MarkBoards(number string, boards *[][][]string, board_state *[][5][5]bool) {
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

func CheckBoards(board_state *[][5][5]bool) int {
	for i, board := range *board_state {
		// Check rows
		for _, row := range board {
			if row == [5]bool{true, true, true, true, true} {
				return i
			}
		}

		// Check cols
		for col := range board[0] {
			for row := range board {
				if board[row][col] == false {
					break
				}

				if row == 4 {
					return i
				}
			}
		}
	}

	return -1
}

func ScoreBoard(board *[5][5]string, board_state *[5][5]bool) (score int) {
	for i, row := range *board_state {
		for j, state := range row {
			if !state {
				number, _ := strconv.Atoi((*board)[i][j])
				score += number
			}
		}
	}

	return
}

func main() {
	input := shared.ReadInputLines("day03/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

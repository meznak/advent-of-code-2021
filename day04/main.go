/*
Advent of Code 2020 day 04
Giant Squid
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

// Find and score a winning bingo card
func SolvePart1(input []string) int {
	var winner int

	numbers := ParseNumbers(input[0])
	boards := ParseBoards(input[2:])

	board_state := make([][5][5]bool, len(boards))

	for _, number := range numbers {
		MarkBoards(number, &boards, &board_state)
		winner = CheckBoards(&board_state, winner)
		if winner > -1 {
			score := ScoreBoard(&boards[winner], &board_state[winner])
			last_num, _ := strconv.Atoi(number)
			return score * last_num
		}
	}

	return -1
}

// Find and score the last winning bingo card
func SolvePart2(input []string) int {
	var winner, win_score int

	numbers := ParseNumbers(input[0])
	boards := ParseBoards(input[2:])
	board_state := make([][5][5]bool, len(boards))

	for _, number := range numbers {
		MarkBoards(number, &boards, &board_state)
		for {
			winner = CheckBoards(&board_state, winner)
			if winner > -1 {
				// Calculate and store winning score
				score := ScoreBoard(&boards[winner], &board_state[winner])
				num_i, _ := strconv.Atoi(number)
				win_score = score * num_i

				// Remove board and its state from the slices
				boards[winner] = boards[len(boards)-1]
				boards[len(boards)-1] = nil
				boards = boards[:len(boards)-1]

				board_state[winner] = board_state[len(board_state)-1]
				board_state = board_state[:len(board_state)-1]
			} else {
				break
			}
		}
	}

	return win_score
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

func CheckBoards(board_state *[][5][5]bool, start int) int {
	if start < 0 {
		start = 0
	}
	for i, board := range (*board_state)[start:] {
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

func ScoreBoard(board *[][]string, board_state *[5][5]bool) (score int) {
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
	input := shared.ReadInputLines("day04/input")
	input = append(input, fmt.Sprintf("\n"))

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

/*
Advent of Code 2020 day 10
Syntax Scoring
*/

package main

import (
	"aoc2021/shared"
	"fmt"
)

// Reject corrupted lines
func SolvePart1(input []string) int {
	score := 0

	// Walk line adding expected closing characters
	// Score and reject line if incorrect closer is detected

	for _, line := range input {
		expected := make([]rune, 0)
		corrupt := false
		depth := 0

		for _, char := range line {
			switch char {
			// openers
			case '(':
				expected = append(expected, ')')
				depth++
			case '[':
				expected = append(expected, ']')
				depth++
			case '{':
				expected = append(expected, '}')
				depth++
			case '<':
				expected = append(expected, '>')
				depth++

				//closers
			case ')':
				if expected[depth-1] == ')' {
					depth--
					expected = expected[:depth]
				} else {
					score += 3
					corrupt = true
				}
			case ']':
				if expected[depth-1] == ']' {
					depth--
					expected = expected[:depth]
				} else {
					score += 57
					corrupt = true
				}
			case '}':
				if expected[depth-1] == '}' {
					depth--
					expected = expected[:depth]
				} else {
					score += 1197
					corrupt = true
				}
			case '>':
				if expected[depth-1] == '>' {
					depth--
					expected = expected[:depth]
				} else {
					score += 25137
					corrupt = true
				}
			}

			if corrupt {
				break
			}
		}
	}

	return score
}

// ???
func SolvePart2(input []string) int {
	return -1
}

func main() {
	input := shared.ReadInputLines("day10/input")

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))
}

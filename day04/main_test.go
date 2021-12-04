package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 4512

	input := shared.ReadInputLines("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart1() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_SolvePart2(t *testing.T) {
	want := -1

	input := shared.ReadInputLines("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart2() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ParseNumbers(t *testing.T) {
	want := []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16", "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"}

	input := shared.ReadInputLines("sample1")

	if got := ParseNumbers(input[0]); !reflect.DeepEqual(got, want) {
		t.Errorf("ParseNumbers() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ParseBoards(t *testing.T) {
	want := [][][]string{
		{
			{"22", "13", "17", "11", "0"},
			{"8", "2", "23", "4", "24"},
			{"21", "9", "14", "16", "7"},
			{"6", "10", "3", "18", "5"},
			{"1", "12", "20", "15", "19"},
		},
		{
			{"3", "15", "0", "2", "22"},
			{"9", "18", "13", "17", "5"},
			{"19", "8", "7", "25", "23"},
			{"20", "11", "10", "24", "4"},
			{"14", "21", "16", "12", "6"},
		},
		{
			{"14", "21", "17", "24", "4"},
			{"10", "16", "15", "9", "19"},
			{"18", "8", "23", "26", "20"},
			{"22", "11", "13", "6", "5"},
			{"2", "0", "12", "3", "7"},
		},
	}

	input := shared.ReadInputLines("sample1")

	if got := ParseBoards(input[2:]); !reflect.DeepEqual(got, want) {
		t.Errorf("ParseBoards() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_MarkBoards(t *testing.T) {
	want := [][5][5]bool{
		{
			{false, true, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
		{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
		{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		},
	}

	input := ParseBoards(shared.ReadInputLines("sample1")[2:])
	got := make([][5][5]bool, len(input))

	if MarkBoards("13", &input, &got); !reflect.DeepEqual(got, want) {
		t.Errorf("MarkBoards() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_CheckBoards_row(t *testing.T) {
	want := 1
	input := [][5][5]bool{
		{
			{false, true, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
		{
			{false, false, false, false, false},
			{true, true, true, true, true},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
		{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		},
	}

	if got := CheckBoards(&input); !reflect.DeepEqual(got, want) {
		t.Errorf("CheckBoards() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_CheckBoards_col(t *testing.T) {
	want := 0
	input := [][5][5]bool{
		{
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
		},
		{
			{false, false, false, false, false},
			{false, true, false, true, true},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
		{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		},
	}

	if got := CheckBoards(&input); !reflect.DeepEqual(got, want) {
		t.Errorf("CheckBoards() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ScoreBoard(t *testing.T) {
	want := 188
	board := [5][5]string{
		{"14", "21", "17", "24", "4"},
		{"10", "16", "15", "9", "19"},
		{"18", "8", "23", "26", "20"},
		{"22", "11", "13", "6", "5"},
		{"2", "0", "12", "3", "7"},
	}
	board_state := [5][5]bool{
		{true, true, true, true, true},
		{false, false, false, true, false},
		{false, false, true, false, false},
		{false, true, false, false, true},
		{true, true, false, false, true},
	}

	if got := ScoreBoard(&board, &board_state); !reflect.DeepEqual(got, want) {
		t.Errorf("ScoreBoard() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Benchmark_SolvePart1(b *testing.B) {
	input := shared.ReadInputLines("input")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SolvePart1(input)
	}
}

func Benchmark_SolvePart2(b *testing.B) {
	input := shared.ReadInputLines("input")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SolvePart2(input)
	}
}

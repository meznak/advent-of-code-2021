package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 5

	input := shared.ReadInputLines("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart1() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_SolvePart2(t *testing.T) {
	want := 12

	input := shared.ReadInputLines("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart2() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ParseLines(t *testing.T) {
	want := [][]int{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
	}

	input := shared.ReadInputLines("sample1")

	if got := ParseLines(input[:4]); !reflect.DeepEqual(got, want) {
		t.Errorf("ParseLines() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ApplyLines_no_diag(t *testing.T) {
	/*
		1 1 3
		. 2 1
		1 2 .
	*/
	want := make(map[Vector]int, 4)
	want[Vector{0, 0}] = 1
	want[Vector{0, 2}] = 1
	want[Vector{1, 0}] = 1
	want[Vector{1, 1}] = 1
	want[Vector{1, 2}] = 2
	want[Vector{2, 0}] = 2
	want[Vector{2, 1}] = 1

	input := [][]int{
		{0, 0, 2, 0},
		{2, 0, 2, 1},
		{1, 2, 0, 2},
		{1, 2, 1, 1},
		{1, 1, 2, 0},
	}

	got := make(map[Vector]int, 0)

	if ApplyLines(&got, &input, true); !reflect.DeepEqual(got, want) {
		t.Errorf("ApplyLines(true) = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ApplyLines_yes_diag(t *testing.T) {
	/*
		1 1 3
		. 2 1
		1 2 .
	*/
	want := make(map[Vector]int, 4)
	want[Vector{0, 0}] = 1
	want[Vector{0, 2}] = 1
	want[Vector{1, 0}] = 1
	want[Vector{1, 1}] = 2
	want[Vector{1, 2}] = 2
	want[Vector{2, 0}] = 3
	want[Vector{2, 1}] = 1

	input := [][]int{
		{0, 0, 2, 0},
		{2, 0, 2, 1},
		{1, 2, 0, 2},
		{1, 2, 1, 1},
		{1, 1, 2, 0},
	}

	got := make(map[Vector]int, 0)

	if ApplyLines(&got, &input, false); !reflect.DeepEqual(got, want) {
		t.Errorf("ApplyLines(false) = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_CountIntersections(t *testing.T) {
	/*
		1 1 2
		. 1 1
		1 2 .
	*/
	input := make(map[Vector]int, 4)
	input[Vector{0, 0}] = 1
	input[Vector{0, 2}] = 1
	input[Vector{1, 0}] = 1
	input[Vector{1, 1}] = 1
	input[Vector{1, 2}] = 2
	input[Vector{2, 1}] = 1
	input[Vector{2, 0}] = 2

	want := 2

	if got := CountIntersections(&input); !reflect.DeepEqual(got, want) {
		t.Errorf("CountIntersections() = %T(%v), want %T(%v)", got, got, want, want)
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

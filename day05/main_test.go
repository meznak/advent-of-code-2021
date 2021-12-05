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
	want := -1

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

func Test_ApplyLines(t *testing.T) {
	/*
		1 1 2
		. 1 1
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
	}

	got := make(map[Vector]int, 0)

	if ApplyLines(&got, &input); !reflect.DeepEqual(got, want) {
		t.Errorf("ApplyLines() = %T(%v), want %T(%v)", got, got, want, want)
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

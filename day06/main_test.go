package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 5934

	input := shared.ReadInputLines("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart1() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_SolvePart2(t *testing.T) {
	want := 26984457539

	input := shared.ReadInputLines("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart2() = %T(%v), want %T(%v)", got, got, want, want)
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

func Test_ProcessDay1(t *testing.T) {
	want := []int{2, 3, 2, 0, 1}

	got := []int{3, 4, 3, 1, 2}

	if ProcessDay(&got); !reflect.DeepEqual(got, want) {
		t.Errorf("ProcessDay(1) = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_ProcessDay2(t *testing.T) {
	want := []int{1, 2, 1, 6, 0, 8}

	got := []int{2, 3, 2, 0, 1}

	if ProcessDay(&got); !reflect.DeepEqual(got, want) {
		t.Errorf("ProcessDay(2) = %T(%v), want %T(%v)", got, got, want, want)
	}
}

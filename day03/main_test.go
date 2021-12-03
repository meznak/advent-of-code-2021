package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 198

	input := shared.ReadInputLines("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart1() = %T(%v), want %T(%v)", got, got, want, want)
	}
}
func Test_SolvePart2(t *testing.T) {
	want := 230

	input := shared.ReadInputLines("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart2() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_MostCommonBit(t *testing.T) {
	want := "10110"

	input := shared.ReadInputLines("sample1")

	var got string

	for i := range input[0] {
		got += string(MostCommonBit(input, i))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MostCommonBit() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_MostCommonBits(t *testing.T) {
	want := "10110"

	input := shared.ReadInputLines("sample1")

	if got := MostCommonBits(input); !reflect.DeepEqual(got, want) {
		t.Errorf("MostCommonBits() = %T(%v), want %T(%v)", got, got, want, want)
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

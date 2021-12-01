package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func test_SolvePart1(t *testing.T) {
	want := 514579

	input := shared.ParseInputInts(shared.ReadInputLines("sample1"))

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %q, want %q", got, want)
	}
}
func test_SolvePart2(t *testing.T) {
	want := 241861950

	input := shared.ParseInputInts(shared.ReadInputLines("sample1"))

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %q, want %q", got, want)
	}
}

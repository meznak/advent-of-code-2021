package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 7

	input := shared.ParseInputInts(shared.ReadInputLines("sample1"))

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}
func Test_SolvePart2(t *testing.T) {
	want := -1

	input := shared.ParseInputInts(shared.ReadInputLines("sample1"))

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}

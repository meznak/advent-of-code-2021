package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 514579

	input := shared.ReadInputInts("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}
func Test_SolvePart2(t *testing.T) {
	want := 241861950

	input := shared.ReadInputInts("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}

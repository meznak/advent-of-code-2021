package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 7

	input := shared.ReadInputInts("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}
func Test_SolvePart2(t *testing.T) {
	want := 5

	input := shared.ReadInputInts("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}

func Test_SumWindows(t *testing.T) {
	want := []int{607, 618, 618, 617, 647, 716, 769, 792}

	input := shared.ReadInputInts("sample1")

	if got := SumWindows(input); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %v, want %v", got, want)
	}
}

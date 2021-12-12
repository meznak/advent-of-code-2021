package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 1656

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

func TestUpdate(t *testing.T) {
	type args struct {
		octos   *[][]int
		flashed *[][]bool
		y       int
		x       int
	}
	tests := []struct {
		name string
		args args
		want *[][]int
	}{
		{
			name: "3x3",
			args: args{
				&[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				&[][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
				0, 0,
			},
			want: &[][]int{
				{2, 3, 4},
				{6, 8, 9},
				{9, 0, 1},
			},
		},
		{
			name: "sample 3x3",
			args: args{
				&[][]int{
					{5, 4, 8},
					{2, 7, 4},
					{5, 2, 6},
				},
				&[][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
				0, 0,
			},
			want: &[][]int{
				{6, 5, 9},
				{3, 8, 5},
				{6, 3, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update(tt.args.octos, tt.args.flashed, tt.args.y, tt.args.x)
			if reflect.DeepEqual(tt.args.octos, tt.want) {
				t.Errorf("SolvePart2() = %T(%v), want %T(%v)", tt.args.octos, tt.args.octos, tt.want, tt.want)
			}
		})
	}
}

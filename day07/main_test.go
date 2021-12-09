package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := []int{2, 37}

	input := shared.ReadInputLines("sample1")

	if got := SolvePart1(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart1() = %T(%v), want %T(%v)", got, got, want, want)
	}
}

func Test_SolvePart2(t *testing.T) {
	want := []int{5, 168}

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

func TestCalculateFuel_Part1(t *testing.T) {
	type args struct {
		crabs *[]int
		fuel  *[][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "four",
			args: args{
				crabs: &[]int{3, 1, 2, 0},
				fuel:  &[][]int{},
			},
			want: [][]int{
				{3, 2, 1, 0},
				{1, 0, 1, 2},
				{2, 1, 0, 1},
				{0, 1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalculateFuel_Part1(tt.args.crabs)
		})
	}
}

func TestCalculateFuel_Part2(t *testing.T) {
	type args struct {
		crabs *[]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "four",
			args: args{
				crabs: &[]int{3, 1, 2, 0},
			},
			want: [][]int{
				{6, 3, 1, 0},
				{1, 0, 1, 3},
				{3, 1, 0, 1},
				{0, 1, 3, 6},
			},
		},
		{
			name: "seven",
			args: args{
				crabs: &[]int{3, 1, 2, 0},
			},
			want: [][]int{
				{6, 3, 1, 0},
				{1, 0, 1, 3},
				{3, 1, 0, 1},
				{0, 1, 3, 6},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFuel_Part2(tt.args.crabs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateFuel_Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFuel(t *testing.T) {
	type args struct {
		fuel *[][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "four",
			args: args{
				&[][]int{
					{3, 2, 1, 0},
					{1, 0, 1, 2},
					{2, 1, 0, 1},
					{0, 1, 2, 3}},
			},
			want: []int{6, 4, 4, 6},
		},
		{
			name: "six",
			args: args{
				&[][]int{
					{3, 2, 1, 0, 3, 5},
					{1, 0, 1, 2, 1, 4},
					{2, 1, 0, 1, 2, 7},
					{0, 1, 2, 3, 0, 2}},
			},
			want: []int{6, 4, 4, 6, 6, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumFuel(tt.args.fuel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	want := 15

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

func TestParseInput(t *testing.T) {
	type args struct {
		input *[]string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{&[]string{"2199943210"}},
			want: [][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLowPoints(t *testing.T) {
	type args struct {
		heights *[][]int
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{
			name: "2x3",
			args: args{
				&[][]int{
					{2, 4, 3},
					{1, 2, 0},
				},
			},
			want: [][]bool{
				{false, false, false},
				{true, false, true},
			},
		},
		{
			name: "5x10",
			args: args{
				&[][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
			},
			want: [][]bool{
				{false, true, false, false, false, false, false, false, false, true},
				{false, false, false, false, false, false, false, false, false, false},
				{false, false, true, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, true, false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLowPoints(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindLowPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcRisk(t *testing.T) {
	type args struct {
		heights *[][]int
		minima  *[][]bool
	}
	tests := []struct {
		name     string
		args     args
		wantRisk int
	}{
		{
			name: "2x3",
			args: args{
				&[][]int{
					{2, 4, 3},
					{1, 2, 0},
				},
				&[][]bool{
					{false, false, false},
					{true, false, true},
				},
			},
			wantRisk: 3,
		},
		{
			name: "3x3",
			args: args{
				&[][]int{
					{2, 4, 1},
					{6, 5, 3},
					{1, 3, 0},
				},
				&[][]bool{
					{true, false, true},
					{false, false, false},
					{true, false, true},
				},
			},
			wantRisk: 8,
		},
		{
			name: "5x10",
			args: args{
				&[][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				&[][]bool{
					{false, true, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, true, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, true, false, false, false},
				},
			},
			wantRisk: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRisk := CalcRisk(tt.args.heights, tt.args.minima); gotRisk != tt.wantRisk {
				t.Errorf("CalcRisk() = %v, want %v", gotRisk, tt.wantRisk)
			}
		})
	}
}

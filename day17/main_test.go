package main

import (
	"aoc2021/shared"
	"reflect"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{
			name: "small",
			args: args{
				shared.ReadInputLines("sample1"),
			},
			want:  []int{6, 9},
			want1: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got1 := SolvePart1(tt.args.input); got[0] != tt.want[0] || got[1] != tt.want[1] || got1 != tt.want1 {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* func Test_SolvePart2(t *testing.T) {
	want := -1

	input := shared.ReadInputLines("sample1")

	if got := SolvePart2(input); !reflect.DeepEqual(got, want) {
		t.Errorf("SolvePart2() = %T(%v), want %T(%v)", got, got, want, want)
	}
} */

func Benchmark_SolvePart1(b *testing.B) {
	input := shared.ReadInputLines("input")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SolvePart1(input)
	}
}

/* func Benchmark_SolvePart2(b *testing.B) {
	input := shared.ReadInputLines("input")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SolvePart2(input)
	}
} */

func TestParseInput(t *testing.T) {
	sample1 := shared.ReadInputLines("sample1")
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "small",
			args: args{
				sample1[0],
			},
			want: []int{20, 30, -10, -5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseInput(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArc(t *testing.T) {
	type args struct {
		probe  Probe
		target []int
	}
	tests := []struct {
		name     string
		args     args
		wantHit  bool
		wantPeak int
	}{
		{
			name: "sample1",
			args: args{
				probe:  Probe{0, 0, 6, 9},
				target: []int{20, 30, -10, -5},
			},
			wantHit:  true,
			wantPeak: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHit, gotPeak := GetArc(tt.args.probe, tt.args.target)
			if gotHit != tt.wantHit {
				t.Errorf("GetArc() gotHit = %v, want %v", gotHit, tt.wantHit)
			}
			if gotPeak != tt.wantPeak {
				t.Errorf("GetArc() gotPeak = %v, want %v", gotPeak, tt.wantPeak)
			}
		})
	}
}

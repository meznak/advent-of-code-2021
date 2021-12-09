package shared

import (
	"testing"
)

func TestAbsInt(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"0", args{0}, 0},
		{"-1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt(tt.args.number); got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		nums *[]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "a",
			args:  args{&[]int{2, 6, 1, 4, 3}},
			want:  2,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinInt(tt.args.nums)
			if got != tt.want {
				t.Errorf("MinInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		nums *[]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "a",
			args:  args{&[]int{2, 6, 1, 4, 3}},
			want:  1,
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxInt(tt.args.nums)
			if got != tt.want {
				t.Errorf("MaxInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	type args struct {
		nums *[]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "4",
			args:    args{&[]int{1, 2, 3, 4}},
			wantSum: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumInt(tt.args.nums); gotSum != tt.wantSum {
				t.Errorf("SumInt() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

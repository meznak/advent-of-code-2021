package shared

import (
	"testing"
)

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

package shared

import (
	"reflect"
	"testing"
)

func TestCSVToInt(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			name:       "test1",
			args:       args{"1,2,3,4,5"},
			wantOutput: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := CSVToInt(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("CSVToInt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

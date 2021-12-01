package shared

import (
	"reflect"
	"testing"
)

func Test_ReadInputLines(t *testing.T) {
	want := []string{"1", "2", "3", "4", "5"}
	if got := ReadInputLines("testdata/test1"); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput() = %q, want %q", got, want)
	}
}

func Test_ReadInputInts(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}

	if got := ReadInputInts("testdata/test1"); !reflect.DeepEqual(got, want) {
		t.Errorf("ParseInputInts() = %q, want %q", got, want)
	}
}

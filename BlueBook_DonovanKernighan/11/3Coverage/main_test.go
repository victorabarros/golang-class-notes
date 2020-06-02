package main

import (
	"testing"
)

func TestMapArrayToString(t *testing.T) {
	attempts := []struct {
		arr  []int
		sep  string
		want string
	}{
		{
			arr:  []int{1, 0, 5},
			sep:  ".",
			want: "1.0.5",
		},
		{
			arr:  []int{2, 9, 5},
			sep:  ";",
			want: "2;9;5",
		},
		{
			arr:  []int{0, 2},
			sep:  ",",
			want: "0,2",
		},
	}
	for _, attempt := range attempts {
		got := MapArrayToString(attempt.arr, attempt.sep)
		if got != attempt.want {
			t.Error(attempt)
		}
	}
}

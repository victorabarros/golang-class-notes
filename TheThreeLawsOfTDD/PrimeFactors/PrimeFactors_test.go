package PrimeFactors

import (
	"reflect"
	"testing"
)

func TestFactorsSuccess(t *testing.T) {
	cases := []struct {
		n   int
		ans []int
	}{
		{
			1,
			[]int{},
		},
		{
			2,
			[]int{2},
		},
		{
			3,
			[]int{3},
		},
		{
			4,
			[]int{2, 2},
		},
		{
			6,
			[]int{2, 3},
		},
		{
			9,
			[]int{3, 3},
		},
		{
			2 * 3 * 5 * 5 * 7 * 7 * 7 * 11,
			[]int{2, 3, 5, 5, 7, 7, 7, 11},
		},
	}
	for _, attempt := range cases {
		if !reflect.DeepEqual(factorsOf(attempt.n), attempt.ans) {
			t.Errorf("Error in attempt %+v\n", attempt)
		}
	}
}

func factorsOf(n int) []int {
	facts := []int{}
	div := 2

	for n > 1 {
		for n%div == 0 {
			facts = append(facts, div)
			n /= div
		}
		div++
	}

	return facts
}

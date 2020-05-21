package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func main() {
	exercise23()
}

func init() {
	for i := range pc {
		// TODO: understand pc[i/2] + byte(i&1)
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func exercise23() {
	// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
	// Compare the performance of the two versions.
	// (Section 11.4 shows how to compare theperformance of different implementations systematically.)
	// fmt.Println(pc)
	var ii uint64 = 0
	for ii < 10 {
		fmt.Println(PopCount(ii), pc[ii], popCountV2(ii))
		ii++
	}
}

func popCountV2(x uint64) int {
	var ans byte = 0
	var ii uint64 = 0
	for ii < 8 {
		ans += pc[byte(x>>(ii*8))]
		ii++
	}
	return int(ans)
}

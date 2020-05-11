package main

// Question: For a constante volume of jobs V, whats behaviour
// of the time of exection with increments threads.
// Let's try:
import (
	"fmt"
	"time"
)

const (
	maxUint uint64 = 1<<64 - 1
	minUint uint64 = 0
)

func main() {
	var target uint64 = minUint
	start := time.Now()

	for target < maxUint {
		target++
	}
	fmt.Println(target, time.Since(start))
}

package main

import (
	"fmt"
	// "math"
	"math/rand"
	"time"
)

// The environment in which these programs are executed is deterministic,
// so each time you run the example program rand.Intn will return the same number.
func hand_rand() {
	var now = time.Now().UnixNano()
	rand.Seed(now)
	for {
		fmt.Println("My favorite number is", rand.Intn(10))
	}
}

func max_int_generator() {
	var max_int int64
	max_int = 8999999999999999999
	for {
		fmt.Println(max_int)
		max_int = max_int * max_int
		// 18446744073709551615
	}
}

func main() {
	// hand_rand()
	max_int_generator()
}

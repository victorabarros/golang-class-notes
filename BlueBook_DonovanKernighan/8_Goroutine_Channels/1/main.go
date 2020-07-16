package main

import (
	"fmt"
	"time"
)

func main() {
	delay := 250 * time.Millisecond
	go spinner(delay)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	time.Sleep(delay * 20)
}

func spinner(delay time.Duration) {
	fmt.Print("Loading\n")
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

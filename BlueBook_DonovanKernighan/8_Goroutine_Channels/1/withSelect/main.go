package main

import (
	"fmt"
	"time"
)

const (
	delay = 100 * time.Millisecond
	n     = 45
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- fib(n)
	}()

	for {
		for _, r := range `-\|/` {
			select {
			case fib := <-ch:
				fmt.Printf("\rFibonacci(%d) = %d\n", n, fib)
				return
			default:
				fmt.Printf("\r%c", r)
				time.Sleep(delay)
			}
		}
	}
	// time.Sleep(time.Duration(delay))
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

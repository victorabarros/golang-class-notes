package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	// unbuffuredChannels41() // Brokes

	// unbuffuredChannels41Success() // Dont brokes
	pipelines42()
}

func unbuffuredChannels41() {
	fmt.Println(<-ch)
}

func unbuffuredChannels41Success() {
	go func() {
		time.Sleep(4 * time.Second)
		ch <- 3
	}()
	// This thread will wait the previous send to ch
	// Unbuffered channels are sometimes called synchronous channels.
	// When a value is sent on an unbuffered channel,
	// the receipt of the value happens before the reawakening of the sending goroutine.
	fmt.Println("Wait send on channel")
	ii, ok := <-ch
	if !ok {
		fmt.Println("not ok")
		return
	}
	fmt.Println(ii)
	fmt.Println("Finished")
}

func pipelines42() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Print("waiting: \t")
		fmt.Print(<-squares, "\n")
	}
}

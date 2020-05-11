package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		fmt.Println("1\t", time.Now().UTC().Sub(now))
		select {
		case <-tick:
			fmt.Println("4\t", time.Now().UTC().Sub(now))
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("5\t", time.Now().UTC().Sub(now))
			fmt.Println("BOOM!")
			// return
			// default:
			//     fmt.Println("2\t", time.Now().UTC().Sub(now))
			//     time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("2\t", time.Now().UTC().Sub(now))
	}
	fmt.Println("10\t", time.Now().UTC().Sub(now))
}

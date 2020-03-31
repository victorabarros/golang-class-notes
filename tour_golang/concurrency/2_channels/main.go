// https://tour.golang.org/concurrency/2
package main

import (
	"fmt"
	// "time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		// fmt.Println(idx, v)
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 2)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// time.Sleep(1 * time.Second)

	fmt.Println(c)
	x, y := <-c, <-c // receive from c
	go sum([]int {1, 2, 3}, c)
	var z int = <-c

	fmt.Println(x, y, x+y, z)
}

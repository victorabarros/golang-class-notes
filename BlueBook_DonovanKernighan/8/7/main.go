package main

import "fmt"

func main() {
	for jj := 0; jj < 5; jj++ {
		poc()
	}
	// example()
}

func example() {
	ch := make(chan int, 2)
	ch <- 100
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("x := <-ch:\t", i) // 1, 3, 5, 7, 9
			fmt.Println("x:\t", x)         // 0 2 4 6 8
		case ch <- i:
			fmt.Println("ch <- i:\t", i) // 0, 2, 4, 6, 8
		default:
			fmt.Println("default") // any print, unless remove chan capacity this way always print default
		}
	}
}

func poc() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 1

	attempts := 10000000.0
	counter1 := 0.0
	counter2 := 0.0

	for ii := 0.0; ii < attempts; ii++ {
		select {
		// The case selected is random
		case <-ch1:
			// fmt.Println("ch1")
			ch1 <- 1
			counter1++
		case <-ch2:
			// fmt.Println("ch2")
			ch2 <- 1
			counter2++
		}
	}
	fmt.Printf("%.1f | %.1f\n", 100*counter1/attempts, 100*counter2/attempts)
}

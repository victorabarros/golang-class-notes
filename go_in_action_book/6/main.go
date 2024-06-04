package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	listing_6_20()
}

var (
	wg sync.WaitGroup
)

func listing_6_20() {
	court := make(chan int)

	wg.Add(2)
	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1
	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s\tWon\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s\tmissed the ball\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s\tHit %d\n", name, ball)

		ball++
		court <- ball
	}
}

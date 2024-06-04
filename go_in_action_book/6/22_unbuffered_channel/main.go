package main

import (
	"fmt"
	"math/rand"
)

var (
	// wg sync.WaitGroup
	court  = make(chan int)
	finish = make(chan int)
)

func main() {
	listing_6_20()
}

func listing_6_20() {

	// wg.Add(2)
	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1
	// wg.Wait()

	// switch <-finish {
	// case 1:
	// 	fmt.Println("Game Over")
	// default:
	// 	fmt.Println("Game Over2")
	// }
	for {
		<-finish
		fmt.Println("Game Over3")
		return
	}
}

func player(name string, court chan int) {
	// defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s\tWon\n", name)
			finish <- 1
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

package main

import (
	"fmt"
	"sync"
)

const (
	nofGoRoutines = 2
	taskLoad      = 4
)

var (
	wg = sync.WaitGroup{}
	// court = make(chan int)
)

func main() {
	listing_6_24()
}

func listing_6_24() {
	tasks := make(chan string, taskLoad)

	for gr := 1; gr <= nofGoRoutines; gr++ {
		wg.Add(1)
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
		// fmt.Println("iter", len(tasks), "/", taskLoad)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		// fmt.Println(ok, "task", task, "/", len(tasks)+1)
		if !ok {
			fmt.Println("Worker", worker, "Shutting\tDown")
			return
		}

		fmt.Println("Worker", worker, "Started\t", task)

		// sleep for a random time to simulate work
		// sleep := rand.Int63n(1)
		// time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("Worker", worker, "Completed\t", task)

	}
}

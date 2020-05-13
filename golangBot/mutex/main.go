// https://golangbot.com/mutex/
package main

import (
	"fmt"
	"sync"
)

func main() {
	incremetWithNoMutex()
	incremetWithInMutex()
}

func incremetWithNoMutex() {
	// The better example of conflict of race condition
	var xx = 0
	var wg sync.WaitGroup

	for ii := 0; ii < 1000; ii++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			xx++
		}()
	}
	wg.Wait()
	fmt.Print(xx, "\t") // 863, 905, 887, 901, 897, 900, 847
}

func incremetWithInMutex() {
	var xx = 0
	var wg sync.WaitGroup
	var m sync.Mutex

	for ii := 0; ii < 1000; ii++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			// If one Goroutine already holds the lock and
			// if a new Goroutine is trying to acquire a lock,
			// the new Goroutine will be blocked until the mutex is unlocked.
			defer m.Unlock()
			xx++
		}()
	}
	wg.Wait()
	fmt.Println(xx) // 1000
}

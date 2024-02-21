// https://pkg.go.dev/sync#Once
package main

import (
	"fmt"
	"sync"
)

func run() {
	var once sync.Once
	max := 10
	done := make(chan bool)
	for i := 0; i < max; i++ {
		go func(ii int) {
			fmt.Println("ff", ii)
			once.Do(func() {
				fmt.Println(ii, "Only once")
			})
			done <- true
		}(i)
	}
	for i := 0; i < max; i++ {
		<-done
	}
}

func run2() {
	var once sync.Once
	max := 10
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()

			fmt.Println("ff", ii)
			once.Do(func() {
				fmt.Println(ii, "Only once")
			})
		}(i)
	}
}

func main() {
	run()
}

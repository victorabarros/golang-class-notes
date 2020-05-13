// https://golang.org/pkg/sync/#Mutex
package main

import (
	"fmt"
	"sync"
	"time"
)

type httpPkg struct{}

func (httpPkg) Get(prefix, url string) {
	fmt.Println(prefix, "\thttp.Get")
	time.Sleep(10 * time.Millisecond)
}

var (
	http httpPkg
	urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
)

func main() {
	// requestsWithGoRoutine() // stdout empty
	requestsWithGoRoutineAndWaitGroup() // Print all urls from slice / dt = 10.26ms
	requests()                          // Print all urls from slice / dt = 750ms
}

func requestsWithGoRoutine() {
	for _, url := range urls {
		go func(url string) {
			http.Get("requestsWithGoRoutine", url)
		}(url)
	}
}

func requestsWithGoRoutineAndWaitGroup() {
	var wg sync.WaitGroup
	// Increment the WaitGroup counter.
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			http.Get("requestsWithGoRoutineAndWaitGroup", url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func requests() {
	for _, url := range urls {
		http.Get("requests", url)
	}
}

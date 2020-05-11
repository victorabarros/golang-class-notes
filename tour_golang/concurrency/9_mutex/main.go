package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	// Lock so only one goroutine at a time can access the map sc.v
	sc.v[key]++
	sc.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	// Lock so only one goroutine at a time can access the map sc.v
	defer sc.mux.Unlock()
	return sc.v[key]
}

func main() {
	now := time.Now()
	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
		// sc.Inc("somekey")
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(sc.Value("somekey"))
	fmt.Println(time.Now().Sub(now))
}

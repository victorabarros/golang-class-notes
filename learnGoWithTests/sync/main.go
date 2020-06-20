package main

import "sync"

func main() {}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()         // With no Mutex the race condition tricks the value.
	defer c.mu.Unlock() // Try comment these lines to see.
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

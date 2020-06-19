package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for time.Since(start) < 100*time.Millisecond {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

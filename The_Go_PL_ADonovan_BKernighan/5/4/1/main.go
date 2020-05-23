package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := WaitForServer("localhost:8080", 30); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for timeout using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string, timeoutSec int) error {
	timeout := time.Duration(timeoutSec) * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 1; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil // success
		}
		log.Printf("attempt %d server not responding (%s); retrying...", tries, err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

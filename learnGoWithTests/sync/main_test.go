package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			t.Run("it runs safely concurrently", func(t *testing.T) {
				wantedCount := 1000
				counter := Counter{}

				var wg sync.WaitGroup
				wg.Add(wantedCount)

				for i := 0; i < wantedCount; i++ {
					go func(w *sync.WaitGroup) {
						counter.Inc()
						w.Done()
					}(&wg)
				}
				wg.Wait()

				if counter.Value() != wantedCount {
					t.Errorf("got %d, want %d", counter.Value(), wantedCount)
				}
			})
		})
}

// Suggestion to look for broken concurrencies: run `go vet`

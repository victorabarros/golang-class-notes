// This sample program demostrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/goinggo/work"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m *namePrinter) Work(id int) {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

func logFunc(message string) {
	log.Println(message)
}

// main is the entry point for all Go programs.
func main() {
	// Create a work value with 2 goroutines.
	w, err := work.New(2, time.Second, logFunc)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specfic name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When Run
				// returns we know it is being handled.
				w.Run(&np)
				wg.Done()
			}()
		}
	}

	for {
		// Enter a number and hit enter to change the size
		// of the work pool.
		var c int
		fmt.Scanf("%d", &c)
		if c == 0 {
			break
		}

		w.Add(c)
	}

	wg.Wait()

	// Shutdown the work and wait for all existing work
	// to be completed.
	w.Shutdown()
}

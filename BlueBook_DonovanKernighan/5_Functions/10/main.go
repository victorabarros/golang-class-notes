package main

import "fmt"

func main() {
	defer func() {
		// Recover catch the panic raised and possibilite handler it.
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
		}
	}()

	letsPanic(3)
}

func letsPanic(ii int) {
	fmt.Printf("letsPanic(%d)\n", ii+0/ii) // panics if ii == 0
	defer fmt.Printf("defer %d\n", ii)

	letsPanic(ii - 1)
}

// https://blog.golang.org/defer-panic-and-recover
package main

import (
	"fmt"
)

func main() {
	// defer it's much usefull to don't duplicate lines where must be invoke
	// in many cases.
	deferExample()
	// behavior of defer in three simple rules:
	deferRule1()
	deferRule2()
	fmt.Println(deferRule3()) // 100 OMG!!!
}

func deferExample() {
	fmt.Println("Starting\n ")
	for ii := 0; ii < 5; ii++ {
		defer fmt.Println(ii)
	}
	fmt.Println("\nFinishing")
}

func deferRule1() {
	// A deferred function's arguments are evaluated
	// when the defer statement is evaluated.
	fmt.Println("\nRule 1:")
	xx := 0
	defer fmt.Println(xx) // 0
	xx++
	fmt.Println(xx) // 1
}

func deferRule2() {
	// Deferred function calls are executed in Last In First Out order
	// after the surrounding function returns.
	fmt.Println("\nRule 2:")
	for ii := 0; ii < 3; ii++ {
		defer fmt.Println(ii)
	}
}

func deferRule3() (jj int) {
	// Deferred functions may read and assign to the returning
	// function's named return values.
	fmt.Println("\nRule 3:")
	defer func() { jj++ }()
	return 99
}

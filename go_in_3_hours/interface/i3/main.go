package main

import "fmt"

func doStuff(i interface{}) {
	switch ii := i.(type) {
	case int:
		fmt.Println("Double i is", ii+ii)
	case string:
		fmt.Println("i is", len(ii), "characters long")
	default:
		fmt.Println("I don't know what to do with this")
	}
}

func main() {
	doStuff(10)
	doStuff("Hello")
	doStuff(true)
}

package main

import (
	"fmt"
	"reflect"
)

// Smartphone its a class abstract from a struct
type Smartphone struct {
	model string
	brand string
	price float32
}

func main() {
	fmt.Println(reflect.TypeOf(Smartphone{}))

	g1 := Smartphone{"galaxy10", "samsung", 754.32}
	fmt.Println(g1)

	s7 := Smartphone{
		model: "S7",
		brand: "apple",
		price: 3950.49,
	}
	fmt.Println(s7)

	// Repricing g1
	g1.price *= 1.12
	fmt.Println(g1)
}

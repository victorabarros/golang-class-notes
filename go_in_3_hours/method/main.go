package main

import (
	"fmt"
)

// Smartphone its a class abstract from a struct
type Smartphone struct {
	model string
	brand string
	price float32
}

// Print with pattern mode
func (s Smartphone) Print() {
	fmt.Printf("Model\t%s\nBrand\t%s\nPrice\tR$%.2f\n\n", s.model, s.brand, s.price)
}

// Off make a discount in percents
func (s *Smartphone) Off(perc float32) {
	s.price *= (1.0 - perc/100)
}

func main() {
	s7 := Smartphone{
		model: "S7",
		brand: "apple",
		price: 5999.99,
	}
	s7.Print()

	s7.Off(10)
	s7.Print()
}

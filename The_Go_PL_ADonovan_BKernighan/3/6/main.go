package main

import "fmt"

type weekday int

const (
	sunday    weekday = iota // 0
	monday                   // 1
	tuesday                  // 2
	wednesday                // 3
	thursday                 // 4
	friday                   // 5
	saturday                 // 6
)

func main() {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
}

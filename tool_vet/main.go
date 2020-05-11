package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) //prints 1
	{
		fmt.Println(x) //prints 1
		x := 2
		fmt.Println(x) //prints 2
	}
	fmt.Println(x) //prints 1 (bad if you need 2)
}

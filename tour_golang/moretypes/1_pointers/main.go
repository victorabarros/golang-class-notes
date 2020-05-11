// https://tour.golang.org/moretypes/1
package main

import (
	"fmt"
	"log"
)

func main() {
	more()
	return
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func more() {
	var a int = 4
	log.Println(&a, a) // address of a, value of a

	b := &a            // point to a : b equals addres of a
	log.Println(b, *b) // address of a, value of a
	log.Println(b == &a)
	log.Println(&b) // addres of pointer b
	// log.Println(*a)  // Breaks! *variable is only vaiable for ponters

	c := a
	log.Println(c, &c)
	log.Println(&c == &a)
	// TODO: how edit lof prefix? show only seconds with decimals
}

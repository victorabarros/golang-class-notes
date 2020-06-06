package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
	fmt.Printf("%T\n", 3)   // "int"

	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"

	t2 := v.Type()           // a reflect.Type
	fmt.Println(t2.String()) // "int"
}

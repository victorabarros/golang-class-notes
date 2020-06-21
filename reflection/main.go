package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
}

func (t *T) Foo() int {
	fmt.Println("foomethod")
	return t.A
}

func main() {
	var t T
	t.A = 5
	r := reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})
	fmt.Println(r)

	s := reflect.ValueOf(&t).Elem()
	a := s.FieldByName("A").Interface()
	fmt.Println(a)
	fmt.Println(a == t.A)
}

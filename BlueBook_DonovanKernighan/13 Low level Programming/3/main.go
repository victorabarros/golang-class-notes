package main

import (
	"fmt"
	"reflect"
	"strings"
)

type mode struct {
	on      bool
	voltage int
}

func main() {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	fmt.Println(got, &got)
	fmt.Println(want, &want)
	fmt.Println(reflect.DeepEqual(got, want))
	// fmt.Println(got == want) ERROR slice can only be compared to nil
	fmt.Println(reflect.DeepEqual([]int{1, 2, 3},
		map[int]bool{1: true, 2: true, 3: true}))
	fmt.Println(reflect.DeepEqual(mode{true, 15}, mode{}))
	fmt.Println(reflect.DeepEqual(mode{true, 15}, mode{true, 30 / 2}))

	a := 3
	b := a
	c := &b
	fmt.Println(a, &a)
	fmt.Println(b, &b)
	fmt.Println(*c, c, &c)
	fmt.Println(a == b)
	fmt.Println(a == *c)
	b++
	fmt.Println(a == *c)
	fmt.Println(*c)
}

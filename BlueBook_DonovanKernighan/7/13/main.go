package main

import "fmt"

type dollar float32
type movie struct {
	title  string
	rating float32
}

func main() {
	wichType(nil)
	wichType(1)
	wichType("str")
	wichType(true)
	wichType(4.5)
	movie1 := movie{"Ghost Buster", 5.0}
	wichType(movie1)
	wichType(movie1.rating)
	jj := struct {
		nome  string
		idade int
	}{
		"Victor Barros", 28,
	}
	wichType(jj)
	wichType(jj.nome)
	ii := dollar(409.43)
	wichType(ii)
	wichType(dollar(133.03))
	wichType([]int{1, 2, 3})
}

func wichType(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println(x, "is type nil")
	case int, uint:
		fmt.Println(x, "is type int, uin")
	case bool:
		fmt.Println(x, "is type bool")
	case string:
		fmt.Println(x, "is type string")
	default:
		fmt.Printf("%v is any order type:\t%T\n", x, x)
	}
}

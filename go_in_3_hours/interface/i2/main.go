package main

import (
    "fmt"
    "reflect"
)

func main() {
    var i interface{}
    fmt.Println("i\t", reflect.TypeOf(i))
    i = "Hello"
    fmt.Println("i\t", reflect.TypeOf(i))
    j, kk := i.(string)
    fmt.Println("j\t", reflect.TypeOf(j), kk)
    k, ok := i.(int)
    fmt.Println("k\t", reflect.TypeOf(k))
    fmt.Println("ok\t", reflect.TypeOf(ok))
    fmt.Println(j, k, ok)
    m, _ := i.(int)
    fmt.Println(m)
    var l interface{}
    fmt.Println("l\t", reflect.TypeOf(l))
    l = 3
    fmt.Println("l\t", reflect.TypeOf(l))
}

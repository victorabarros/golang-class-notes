package main

import (
    "fmt"
    "reflect"
)

func main() {
    m := make(map[string]int)
    m["hello"] = 300
    h := m["hello"]
    fmt.Println("hello in m:", h)
    fmt.Println("a in m:", m["a"])

    if v, ok := m["hello"]; ok {
        fmt.Println("hello in m:", v)
        fmt.Println(ok)
        fmt.Println(reflect.TypeOf(v), reflect.TypeOf(ok))
    }

    m["hello"] = 20
    fmt.Println("hello in m:", m["hello"])

    m2 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 50,
    }

    for k, v := range m2 {
        fmt.Println(k, v)
    }

    var slice []int8
    var arr [2]int8
    var _map map[string]int
    fmt.Println(slice, arr, _map)
}

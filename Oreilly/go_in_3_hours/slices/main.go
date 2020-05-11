package main

import (
    "fmt"
    "reflect"
)

func main() {
    var s1 = [2]int{1, 2}

    s2 := s1[:1]
    fmt.Println(s1, s2)

    s2[0] = 50
    fmt.Println(s1, s2)
    fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(s2))

    s1[0] = 55
    fmt.Println(s1, s2)

    modArray(s1)
    fmt.Println(s1, s2)

    modSlice(s2)
    fmt.Println(s1, s2)

    uniHello := "👋 🌍"
    bytes := []byte(uniHello)
    fmt.Println(bytes)
    runes := []rune(uniHello)
    runes[1] = 'a'

    for _, v := range runes {
        fmt.Println(string(v))
        fmt.Println(reflect.TypeOf(v))
    }
}

func modSlice(s []int) {
    s[0] = 44
}

func modArray(s [2]int) {
    s[0] = 33
}

package main

import "fmt"

func main() {
    for {
        example11()
    }
}

func example11() {
    var data int
    go func() {
        data++
    }()
    if data == 0 {
        fmt.Printf("the value is %v.\n", data)
    } else {
        fmt.Print("Data not Zero.\n")
    }
}

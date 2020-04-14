// https://tour.golang.org/concurrency/4
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        fmt.Println("c<-x")
        c <- x
        x, y = y, x+y
    }
    // Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
    // close(c)
}

func main() {
    c := make(chan int)
    go fibonacci(15, c)
    fmt.Println("1 iter")
    for i := range c {
        fmt.Println(i)
    }

    go fibonacci(cap(c), c)
    fmt.Println("2 iter")  // Uncomment this line the code breaks. TODO: NEED HELP.
    for range c{
        fib, ok := <-c
        fmt.Println(fib, ok)
    }

    fmt.Println("~~~~~~~~")
}

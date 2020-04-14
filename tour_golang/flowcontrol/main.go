// https://tour.golang.org/flowcontrol/2
// https://tour.golang.org/flowcontrol/3
package main

import "fmt"

func main() {
    sum := 1
    // for working as while
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}

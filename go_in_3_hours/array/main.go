package main

import "fmt"

func main() {
    var vals1 [2]int8
    vals1[0] = 0
    vals1[1] = 1

    fmt.Println("vals1 =", vals1)

    fmt.Println("vals1[0] + 1")
    vals1[0]++
    fmt.Println("vals1 =", vals1)

    var vals2 [2]int8
    fmt.Println("\nvals2 =", vals2)
    vals2[0] = 5
    fmt.Println("vals2 =", vals2)
    // fmt.Println("vals1 + vals2 =", vals1+vals2)
}

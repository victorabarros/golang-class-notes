package main

import "fmt"

func main() {
    var ii int8 = 10
    fmt.Println("ii =\t", ii)
    var jj float32 = 10 - 0.00001
    fmt.Println("jj =\t", jj)
    var kk float64 = 10.01
    fmt.Println("kk =\t", kk)
    fmt.Println("ii + jj =\t", float32(ii)+jj)
}

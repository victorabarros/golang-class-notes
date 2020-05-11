package main

import "fmt"

func main() {
    fmt.Println("\nLearning if:")
    var ii byte = 0
    if jj := ii / 2; jj != 0 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

    fmt.Println("\n\nLearning for kk := 0; kk < 10; kk++ :")
    for kk := 0; kk < 10; kk++ {
        if kk == 3 {
            //Não printa 3
            continue
        } else if kk == 5 {
            //Interrompe
            break
        }
        fmt.Println(kk)
    }

    fmt.Println("\n\nLearning for ll < 3:")
    ll := 0
    for ll < 3 {
        fmt.Println(ll)
        ll++
    }

    fmt.Println("\n\nLearning for:")
    mm := 0
    for {
        fmt.Println(mm)
        mm++
        if mm == 4 {
            break
        }
    }

    fmt.Println("\n\nLearning for range string:")
    nn := "VictorBarros"
    for k, v := range nn {
        fmt.Println(k, "\t", v, "\t", string(v))
    }
}

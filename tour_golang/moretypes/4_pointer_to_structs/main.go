// https://tour.golang.org/moretypes/4
package main

import (
    "fmt"
    "log"
)

type Vertex struct {
    X int
    Y int
}

func main() {
    more()
    return
    v := Vertex{11, 12}

    p := &v
    p.X = 13
    fmt.Println(v)
}

func more() {
    v := Vertex{11, 12}
    log.Println("v", "\t", &v, "\t", v)

    p := &v
    log.Println("p", "\t", &p, "\t", p)
    p.X = 13
    log.Println("v", "\t", &v, "\t", v, "\t", *&v)

    w := v
    log.Println("w", "\t", &w, "\t", w, "\t", *&w)
    w.X = 14

    v.Y = 15
    log.Println("w", "\t", &w, "\t", w, "\t", *&w)
    log.Println("v", "\t", &v, "\t", v, "\t", *&v)
}

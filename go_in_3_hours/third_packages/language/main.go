package main

import (
    "fmt"

    "github.com/victorabarros/home/victor/Documents/repository/Learning/golang/go_in_3_hours/third_packages/language/mapper/mapper.go"
)

func main() {
    fmt.Println(mapper.Greet("Howdy, what's new?"))
    fmt.Println(mapper.Greet("Comment allez vous?"))
    fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}

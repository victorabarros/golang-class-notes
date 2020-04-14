package main

import (
    "fmt"

    "home/victor/Documents/repository/Learning/golang/go_in_3_hours/packages/creating/leftpad"
)

func main() {
    fmt.Println(leftpad.Format("hello", 15))
    fmt.Println(leftpad.Format("goodbye", 15))
    fmt.Println(leftpad.Format("¿Cómo está?", 15))
    fmt.Println(leftpad.Format("Internationalization", 15))
    fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
    fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
    fmt.Println(leftpad.FormatRune("¿Cómo está?", 15, '😀'))
}

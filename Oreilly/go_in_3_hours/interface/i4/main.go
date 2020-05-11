package main

import "fmt"

type tester interface {
    test(int) bool
}

func runTests(i int, tests []tester) bool {
    result := true
    for _, _test := range tests {
        result = result && _test.test(i)
    }
    return result
}

type testerFunc func(int) bool

func (tf testerFunc) test(i int) bool {
    fmt.Println("---1")
    return tf(i)
}

func main() {
    result := runTests(11, []tester{
        testerFunc(func(i int) bool {
            fmt.Println("---2")
            return i%2 == 0
        }),
        testerFunc(func(i int) bool {
            fmt.Println("---3")
            return i < 20
        }),
    })
    fmt.Println(result)
}

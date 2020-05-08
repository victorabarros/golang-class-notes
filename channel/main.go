package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	dict := make(map[int]bool)
	go send(ch, dict)
	ii := 0
	for {
        fmt.Println("1")
		ii++
		time.Sleep(250 * time.Millisecond)
		dict[ii] = true
		ch <- ii
	}
}

func send(ch chan int, dict map[int]bool) {
	for {
        fmt.Println("2")
		time.Sleep(1*time.Second)
		fmt.Println(mapp(ch, dict))
    }
    dict = make(map[int]bool)
}

func mapp(ch chan int, dict map[int]bool) []int{
	result := []int{}
    fmt.Println("3")
	for _ = range dict {
        fmt.Println("4")
		result = append(result, <-ch)
        fmt.Println(result)
	}
	return result
}

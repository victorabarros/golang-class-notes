// Goroutines
// https://youtu.be/YS4e4q9oBaU?t=20922
package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var now = time.Now().UTC()

func main() {
	for i:= 0; i < 10; i++{
		wg.Add(2)
		go sayHello(i)
		go increment()
	}
	wg.Wait()
}

func sayHello(i int){
	var dt = time.Now().UTC().Sub(now)
	fmt.Println(i, "\t", dt, "\t", counter)
	wg.Done()
}

func increment(){
	counter ++
	wg.Done()
}
// Continue with https://youtu.be/YS4e4q9oBaU?t=20906

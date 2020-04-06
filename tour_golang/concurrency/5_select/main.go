// https://tour.golang.org/concurrency/5
package main

import (
	"fmt"
	"time"
)

var NOW =  time.Now()

func fibonacci(cc, quit chan int) {
	fmt.Println(time.Now().Sub(NOW), "\t2 fibo start")
	xx, yy := 0, 1
	var sum int
	for {
		select {
		case cc<-xx:
			fmt.Println(time.Now().Sub(NOW), "\tcc<-xx")
			sum = xx + yy
			xx = yy
			yy = sum
		case <-quit:
			fmt.Println(time.Now().Sub(NOW), "\tquit")
			return
		}
	}
}

func main() {
	fmt.Println("---\n", time.Now().Sub(NOW), "\t\t1 main start")
	cc := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println(time.Now().Sub(NOW), "\t3 func start")
		fmt.Println(time.Now().Sub(NOW), "\t<-cc [1]\t",<-cc)
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now().Sub(NOW), "\t<-cc [2]\t",<-cc)
		}
		fmt.Println(time.Now().Sub(NOW), "\t<-cc [3]\t",<-cc)
		fmt.Println(time.Now().Sub(NOW), "\tquit<-0\t")
		quit<-0
	}()
	fibonacci(cc, quit)
	fmt.Println("---")
}

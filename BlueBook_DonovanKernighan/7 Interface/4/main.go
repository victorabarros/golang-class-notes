package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	now := time.Now()

	fmt.Println(time.Now().Sub(now), "\t1")
	flag.Parse()
	fmt.Printf("%s\tSleeping for %v...\n", time.Now().Sub(now), *period)
	time.Sleep(*period)
	fmt.Println(time.Now().Sub(now), "\t2")
}

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

type Storage struct {
	sync.RWMutex
	HashMap map[string]string
}

var storage = Storage{
	HashMap: make(map[string]string),
}

func runApp(numOfTimes int, c chan<- struct{}) {
	id := fmt.Sprintf("%s", time.Now().Format("20060102150405"))

	for t := 1; t <= numOfTimes; t++ {
		storage.Lock()
		storage.HashMap[id] = id
		storage.Unlock()
	}

	c <- struct{}{}
}

func main() {
	fmt.Println("Starting...")

	cpuProfile, _ := os.Create("./report/cpuprofile")
	memProfile, _ := os.Create("./report/memprofile")

	pprof.StartCPUProfile(cpuProfile)

	fmt.Println("Started")

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go runApp(50000, c1)
	go runApp(50000, c2)
	go runApp(50000, c3)

	<-c1
	<-c2
	<-c3

	pprof.StopCPUProfile()
	pprof.WriteHeapProfile(memProfile)

	fmt.Println("Finished!!!")
}

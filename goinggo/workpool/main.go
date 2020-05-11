package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// from https://github.com/goinggo/workpool/blob/master/workpool.go
type WorkPool struct {
	shutdownQueueChannel chan string
	shutdownWorkChannel  chan struct{}
	shutdownWaitGroup    sync.WaitGroup
	queueChannel         chan poolWork
	workChannel          chan PoolWorker
	queuedWork           int32
	activeRoutines       int32
	queueCapacity        int32
}

type poolWork struct {
	Work          PoolWorker
	ResultChannel chan error
}

type PoolWorker interface {
	DoWork(workRoutine int)
}

// Duas structs e uma interface todos com o mesmo fucking nome.
// Legibilidate e criatividade ZERO

type MyTask struct {
	Name string
	WP   *WorkPool
}

func (mt *MyTask) DoWork(workRoutine int) {
	fmt.Println(mt.Name)

	fmt.Printf("*******> WR: %d QW: %d AR: %d\n",
		workRoutine,
		mt.WP.QueuedWork(),
		mt.WP.ActiveRoutines())

	time.Sleep(100 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	workPool := workpool.New(runtime.NumCPU()*3, 100)

	task := MyTask{
		Name: "A" + strconv.Itoa(i),
		WP:   workPool,
	}

	err := workPool.PostWork("main", &task)

	// …
}

func New(numberOfRoutines int, queueCapacity int32) *WorkPool {
	workPool := WorkPool{
		shutdownQueueChannel: make(chan string),
		shutdownWorkChannel:  make(chan struct{}),
		queueChannel:         make(chan poolWork),
		workChannel:          make(chan PoolWorker, queueCapacity),
		queuedWork:           0,
		activeRoutines:       0,
		queueCapacity:        queueCapacity,
	}

	for workRoutine := 0; workRoutine < numberOfRoutines; workRoutine++ {
		workPool.shutdownWaitGroup.Add(1)
		go workPool.workRoutine(workRoutine)
	}

	go workPool.queueRoutine()
	return &workPool
}

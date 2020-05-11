package main

import (
    "fmt"
    "time"
    "strconv"
)

// from https://github.com/goinggo/workpool/blob/master/workpool.go
// type WorkPool struct {
//     shutdownQueueChannel chan string
//     shutdownWorkChannel  chan struct{}
//     shutdownWaitGroup    sync.WaitGroup
//     queueChannel         chan poolWork
//     workChannel          chan PoolWorker
//     queuedWork           int32
//     activeRoutines       int32
//     queueCapacity        int32
// }
// type poolWork struct {
//     Work          PoolWorker
//     ResultChannel chan error
// }
// type PoolWorker interface {
//     DoWork(workRoutine int)
// }

// Duas structs e uma interface todos com o mesmo fucking nome.
// Legibilidate e criatividade ZERO
type MyTask struct {
    Name string
    WP *workpool.WorkPool
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

    workPool := workpool.New(runtime.NumCPU() * 3, 100)

    task := MyTask{
        Name: "A" + strconv.Itoa(i),
        WP: workPool,
    }

    err := workPool.PostWork("main", &task)

    // …
}

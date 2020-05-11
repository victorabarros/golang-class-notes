package workpool

import "sync"

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

package main

// Question: For a constante volume of jobs V, whats behaviour
// of the time of exection with increments threads.
// Let's try:

import (
	"fmt"
	"time"
)

const (
	maxInt   int           = 1<<24 - 1
	minInt   int           = 0
	nThreads int           = 256
	sleepNs  time.Duration = time.Duration(2*1e9/maxInt) * time.Nanosecond
)

var (
	value int = minInt
)

func main() {
	start := time.Now()

	for ii := 0; ii < nThreads; ii++ {
		go incrementUntilTargetAndSleep()
	}
	for value < maxInt {
		// wait
	}
    fmt.Println("dt: ", time.Since(start))
    // TODO: use Benchmark. Example in ./Oreilly/The_Go_Programming_Language_by_Alan_A_A_Donovan_and_Brian_W_Kernighan/11/4Benchmark
}

func incrementUntilTarget() {
	for value < maxInt {
		value++
		// fmt.Println(value, maxInt)
	}
}

// Result incrementUntilTarget with maxInt = 1<<32 - 1
// nThreads = 1: dt = [5.767828906, 5.747121803, 5.80211074, 5.789619506, 5.775224127, 5.779455311, 5.798811127]
// nThreads = 2: dt = [9.127142756, 11.09388874, 11.616361915, 9.539769533, 9.476375706]
// nThreads = 32: dt = [20.237369295]

func incrementUntilTargetAndSleep() {
	for value < maxInt {
		value++
		time.Sleep(sleepNs)
		// fmt.Println(value, maxInt)
	}
}

// Result incrementUntilTargetAndSleep with maxInt = 1<<24 - 1 and sleepNs = time.Duration(2*1e9/maxInt) * time.Nanosecond
// nThreads =   1: dt = [39.270725689, 25.919333473, 26.776576502, 18.686939505]
// nThreads =   2: dt = [6.355238534, 6.258431474, 6.264812475, 5.494026552, 5.424849386, 5.6991336, 6.120461255]
// nThreads =   4: dt = [2.514402943, 2.822076708, 2.800692265, 2.612810489, 2.557452482, 2.972176422, 2.416671531]
// nThreads =   8: dt = [1.741156632, 1.716172212, 1.715971609, 1.693476771, 1.757831546, 1.795138854, 1.879073632]
// nThreads =  16: dt = [1.691664248, 1.897603846, 1.689581334, 1.703788956, 1.754558852, 1.749904504, 1.871445892, 1.856035925]
// nThreads =  32: dt = [1.724537122, 1.707977846, 1.711238373, 1.756726488, 1.794056966, 1.870512965, 1.844033489]
// nThreads =  64: dt = [1.713174181, 1.72437667, 1.703943977, 1.709811021, 1.695212623, 1.687716622, 1.75760317]
// nThreads = 128: dt = [1.711989813, 1.721393932, 1.706521077, 1.697823659, 1.785392965, 1.775768778, 1.790022245]
// nThreads = 256: dt = [1.731741538, 1.70834228, 1.697091836, 1.673216636, 1.755612303, 1.807154102]

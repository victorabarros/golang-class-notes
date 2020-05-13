package main

import (
	"fmt"
	"time"
)

func main() {
	Idx := make(map[int]bool)
	bufferCh := make(chan int)
	go BufferHandler(&Idx, bufferCh)
	for ii := 0; ; ii++ {
		time.Sleep(time.Duration(3) * time.Second)
		fmt.Printf("\nnew message")
		SKU := ii
		if _, ok := Idx[SKU]; ok {
			fmt.Printf("\n%d Already on buffer.", SKU)
			continue
		}

		Idx[SKU] = true
		bufferCh <- SKU
	}
}

// BufferHandler wait the timeout, makes request to OTA in batch and clean buffer
func BufferHandler(
	Idx *map[int]bool, bufferCh chan int) {
	timeout := 10
	for {
		quit := make(chan bool)
		go func() {
			time.Sleep(time.Duration(timeout) * time.Second)
			*Idx = make(map[int]bool)
			quit <- true
		}()

		SKUArray := mapSKUChanToArray(bufferCh, quit)

		go func() {
			for SKU := range SKUArray {
				fmt.Println("cfg.OTA.URL", SKU)
				time.Sleep(time.Duration(timeout/len(SKUArray)) * time.Second)
			}
		}()
		fmt.Printf("\n%d request to %s", len(SKUArray), "cfg.OTA.URL")
	}
}

// mapSKUChanToArray build a array from a channel
func mapSKUChanToArray(origin chan int, quit chan bool) []int {
	SKUs := []int{}
	for {
		select {
		case SKU := <-origin:
			SKUs = append(SKUs, SKU)
		case <-quit:
			return SKUs
		}
	}
}

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	time.Sleep(time.Second)
	c := make(chan bool)
	<-c
}

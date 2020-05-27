package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	printColor = "\033[38;5;%dm%s\033[39;49m"
)

func main() {
	logrus.Info("Starting script ~logging with logrus pkg")
	for j := 0; j < 256; j++ {
		fmt.Printf(printColor, j, "victorbarros1130! ")
	}
	logrus.Info("Finishing")
}

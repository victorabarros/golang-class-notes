package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateString(length int) string {
	rand.Seed(time.Now().Unix())
	ba := make([]byte, length)
	rand.Read(ba)
	randomStr := fmt.Sprintf("%x", ba)
	return randomStr[:length]
}

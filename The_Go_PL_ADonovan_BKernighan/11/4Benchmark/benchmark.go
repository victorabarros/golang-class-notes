package benchmark

import (
	"fmt"
	"strings"
	"time"
)

// mapArrayToString turn any array of integer into a string
func mapArrayToString(arr []int, sep string) string {
	// Font: https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string
	return strings.Trim(
		strings.Join(strings.Fields(fmt.Sprint(arr)), sep),
		"[]",
	)
}

func nMillisecondsSleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func nNanosecondsSleep(n int) {
	time.Sleep(time.Duration(n) * time.Nanosecond)
}

func isPalindromeV1(word string) bool {
	length := len(word)
	for ii := 0; ii < length; ii++ {
		if word[ii] != word[length-ii-1] {
			return false
		}
	}
	return true
}

func isPalindromeV2(word string) bool {
	length := len(word)
	for ii := 0; ii <= length/2; ii++ {
		if word[ii] != word[length-ii-1] {
			return false
		}
	}
	return true
}

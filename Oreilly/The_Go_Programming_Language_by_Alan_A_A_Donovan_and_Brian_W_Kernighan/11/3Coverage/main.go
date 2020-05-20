package main

import (
	"fmt"
	"strings"
)

func main() {
}

// MapArrayToString turn any array of integer into a string
func MapArrayToString(arr []int, sep string) string {
	// Font: https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string
	return strings.Trim(
		strings.Join(strings.Fields(fmt.Sprint(arr)), sep),
		"[]",
	)
}

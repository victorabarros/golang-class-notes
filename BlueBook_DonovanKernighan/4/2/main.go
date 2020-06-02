package main

import "fmt"

func main() {
	s := [6]int{0, 1, 2, 3, 4, 5} // Rotate s left by two positions.
	fmt.Println(s)
	reverse(s[:2])
	reverse(s[2:])
	reverse(s[:])
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

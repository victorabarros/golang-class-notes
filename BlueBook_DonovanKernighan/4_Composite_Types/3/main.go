package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings // a kind of cache

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("insert an value")
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
		fmt.Println("")
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]

	// more complex switch:
	if word == "complex" {
		word2 := os.Args[2]
		fmt.Println("\n" + word2)
		c := "crackerjack"
		switch l := len(word2); {
		case word2 == "hi":
			fmt.Println("informal")
			fallthrough
		case word2 == "hello":
			fmt.Println("How R you?")
		case l == 1:
			fmt.Println("Doesn't exist word with one letter.")
		case l > 5 && l < 10, word2 == c:
			fmt.Println("Word 6-9 characters long or eihter", c)
		default:
			fmt.Println("Don't know what to say.")
		}
	} else {
		fmt.Println("\n" + word)
		switch word {
		case "hello":
			fmt.Println("Hi man")
			fallthrough
		case "greet":
			fmt.Println("Salutations")
		case "bye", "goodbye":
			fmt.Println("So long")
		default:
			fmt.Println("Sorry?")
		}
	}
}

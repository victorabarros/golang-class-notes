package main

// Findlinks1 prints the links in an HTML document read from standard input.package main
import (
	"fmt"
	"os"
	// "golang.org/x/net/html"
)

func main() {
	// doc, err := html.Parse(os.Stdin)
	doc, err := []string{}, fmt.Errorf("", nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	// Change the findlinks program to traverse then.FirstChild linked list using recursive calls to visitinstead of a loop.
	exercise1(doc)
}

func exercise1(doc []string) {
	if len(doc) == 0 {
		return
	}

	fmt.Println(doc[0])
	exercise1(doc[:len(doc)-1])
}

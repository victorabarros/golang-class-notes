package main

import (
	"fmt"

	"github.com/victorabarros/golang-class-notes/go_in_3_hours/third_packages/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}

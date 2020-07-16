package main

import "fmt"

func main() {
	q := [...]int{1, 2, 3} // length do tamanho da quantidade de parametros inseridos na sequencia
	fmt.Printf("%T\n", q)  // "[3]int"
	r := [...]int{99: -1}
	fmt.Print(r, len(r))
}

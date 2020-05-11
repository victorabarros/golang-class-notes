package main

import "fmt"

func main() {
	fmt.Println("\n addNumbers")
	addNumbers(3, 4)
	addNumbers(30, -4)
	addNumbers(30, -100)

	fmt.Println("\n sum")
	fmt.Println(sum(30, -4))
	fmt.Println(sum(30, -100))

	fmt.Println("\n divAndRemainder")
	fmt.Println(divAndRemainder(3, 4))
	fmt.Println(divAndRemainder(100, 30))

	fmt.Println("\n doubleInside")
	var a int8 = 1
	arr, s := [2]int8{2, 3}, "hello"
	fmt.Println(a, arr, s)
	doubeOnlyInside(a, arr, s)
	fmt.Println(a, arr, s)
}

func addNumbers(a int8, b int8) {
	fmt.Println(a + b)
}

func sum(a int8, b int8) int8 {
	return a + b
}

func divAndRemainder(a int8, b int8) (int8, int8) {
	//divisão e resto
	return a / b, a % b
}

func doubeOnlyInside(a int8, arr [2]int8, s string) {
	// Operações algébricas em viariáveis dos parÂmetros não alteram seu valor global.
	a *= 2
	for ii := 0; ii < len(arr); ii++ {
		arr[ii] *= 2
	}
	s += s
	fmt.Println(a, arr, s)
}

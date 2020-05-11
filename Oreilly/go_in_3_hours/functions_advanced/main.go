package main

import "fmt"

func globalSum(a int, b int) int {
	return a + b
}

func printOperation(a int, b int, f func(int, int) int) {
	fmt.Println("l1\t", f(a, b))
}

func makerAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func main() {
	localGlobalSum := globalSum
	fmt.Println("l2\t", localGlobalSum(5, 2))
	printOperation(5, 2, globalSum)
	printOperation(5, 2, localGlobalSum)

	c := 2
	localSumC := func(a int) int {
		c += a // Change C value
		return c
	}
	fmt.Println("l3\t", localSumC(2))
	fmt.Println("l4\t", localSumC(2))

	localDouble := func(a int) int {
		a *= 2
		return a
	}
	fmt.Println("l5\t", localDouble(2))
	fmt.Println("l6\t", localDouble(2))

	addOne := makerAdder(1)
	fmt.Println("l7\t", addOne(14))
}

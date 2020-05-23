package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"

	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	// Exercise 5.15: Write variadic functions max and min, analogous to sum.
	// What should these functions do when called with no arguments?
	fmt.Println("\nExercise 5.15")
	fmt.Println(max(values...))      // 4 <nil>
	fmt.Println(max(1, 2, 3, 4, 50)) // 50 <nil>
	fmt.Println(max())               // 0 Input params empty

	// Exercise 5.16: Write a variadic version of strings.Join.
	fmt.Println("\nExercise 5.16")
	fmt.Println(join("1", "2", "3", "4", "50"))              // 21314150
	fmt.Println(join(",", []string{"2", "3", "4", "50"}...)) // 2,3,4,50

	fmt.Println("\nExercise 5.16 V2")
	fmt.Println(joinV2(",", "2", 3, true)) // 21314150
}

func sum(vals ...int) int {
	// fmt.Printf("%T\n", vals) // "[]int"
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("Input params empty")
	}

	max := math.MinInt64
	for _, ii := range vals {
		if ii > max {
			max = ii
		}
	}
	return max, nil
}

func join(sep string, vals ...string) (ans string) {
	ans = vals[0]
	for _, ii := range vals[1:] {
		ans += sep + ii
	}
	return
}

func joinV2(sep interface{}, vals ...interface{}) (ans string) {
	// fmt.Printf("%T\n", vals) // "[]interface {}"
	ans = fmt.Sprintf("%s", vals[0])
	for _, ii := range vals[1:] {
		// TODO: Melhorar com o % correto de acordo com o type.
		// Fazer um switch case com o type para determinar %s, %d, %b...
		ans += fmt.Sprintf("%s%s", sep, ii)
	}
	return
}

package main

import (
	"fmt"
	"reflect"
)

func setTo10(aa *int) {
	*aa = 10
}

func setTo20Fail(bb *int) {
	fmt.Println(*bb, "\t", bb)
	bb = new(int)
	fmt.Println(*bb, "\t", bb)
	*bb = 20
	fmt.Println(*bb, "\t", bb)
}

func main() {
	a := 1
	b := &a // b aponta para o endereço de memória de a. Logo *b == a. &a é um pointer/reference de a
	c := a  // c tem é igual à a instantaneamente, mas não apontam para o mesmo endereço de memória
	fmt.Println("\nInitialized", "\na\t", a, "\n&a\t", &a, "\nb\t", b, "\n*b\t", *b, "\n&b\t", &b, "\nc\t", c, "\n&c\t", &c)

	a = 5
	fmt.Println("\nSetting a = 5", "\na\t", a, "\n&a\t", &a, "\nb\t", b, "\n*b\t", *b, "\n&b\t", &b, "\nc\t", c, "\n&c\t", &c)

	*b = 7 // Para resgatar o valor do endereço de memória usa-se o * antes da variável
	fmt.Println("\nSetting *b = 7", "\na\t", a, "\n&a\t", &a, "\nb\t", b, "\n*b\t", *b, "\n&b\t", &b, "\nc\t", c, "\n&c\t", &c)

	c = 3
	fmt.Println("\nSetting c = 3", "\na\t", a, "\n&a\t", &a, "\nb\t", b, "\n*b\t", *b, "\n&b\t", &b, "\nc\t", c, "\n&c\t", &c)

	fmt.Println("\n", *&a, reflect.TypeOf(b), b, &b, &a) // Tests

	fmt.Println("\nDeclaring a variable as a pointer")
	var d *int

	fmt.Println(d)
	// fmt.Println(*d) // PANIC
	fmt.Println(&d)

	fmt.Println("\nOther way to")
	e := new(int)
	fmt.Println(e)
	fmt.Println(*e)
	fmt.Println(&e)

	fmt.Println("\nIt's possible change the variable value calling a function out of main")
	f := 9
	fmt.Println(&f, "before function\t", f)
	setTo10(&f)
	fmt.Println(&f, "after function\t", f)

	fmt.Println("\nNOTICE:\nGO it's a call-by-value language.\nDiferente from python, for example.")

	fmt.Println("\nNow if we assign a new pointer inside of the function, just add one new line.")
	g := 13
	fmt.Println(&g, "before function\t", g)
	setTo20Fail(&g)
	fmt.Println(&g, "after function\t", g)
	fmt.Println("No changes.\nA linha \"new(int)\" não resetou o valor para 0 como deu um novo endereço de memória.\nEntão a variável que foi atribuida valor de 20 não é mais a mesma que veio do parâmetro.")
}

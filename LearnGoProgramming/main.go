package main

import "fmt"

type product interface {
	print()
}

type book struct {
	Name string
}

func (b book) print() { fmt.Println("Book\t", b.Name) }

type game struct {
	Name string
}

func (g game) print() { fmt.Println("Game\t", g.Name) }

func main() {
	stock := []product{
		book{"Moby Dick"},
		game{"GTA"},
	}

	for _, ii := range stock {
		ii.print()
	}
}

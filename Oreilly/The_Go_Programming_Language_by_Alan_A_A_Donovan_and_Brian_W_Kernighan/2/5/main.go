package main

// Package tempconv performs Celsius and Fahrenheit temperature computations.package tempconv
import "fmt"

type Celsius float64
type Fahrenheit float64

var (
	absoluteZeroC Celsius = -273.15
	freezingC     Celsius = 0
	boilingC      Celsius = 100
)

func main() {
	tempCelsius := Celsius(45.7)
	fmt.Println(CToF(tempCelsius))
	fmt.Println(FToC(Fahrenheit(114.26)))

	fmt.Println(CToF(absoluteZeroC))
	fmt.Println(CToF(freezingC))
	fmt.Println(CToF(boilingC))

	// Don't you think it's better code
	// A method func (c *Celsius) ToFahrenheit Fahrenheit {...}
	// And  func (c *Fahrenheit) ToCelsius Celsius {...}
	// ?
	fmt.Println("\nV2:")
	mainV2()
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c *Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(*c*9/5 + 32) }
func (f *Fahrenheit) ToCelsius() Celsius    { return Celsius((*f - 32) * 5 / 9) }

func mainV2() {
	tCelsius := Celsius(45.7)
	fmt.Println(tCelsius.ToFahrenheit())
	tFahrenheit := Fahrenheit(114.26)
	fmt.Println(tFahrenheit.ToCelsius())
	// fmt.Println(Fahrenheit(114.26).ToCelsius()) // Why don't work???

	fmt.Println(absoluteZeroC.ToFahrenheit())
	fmt.Println(freezingC.ToFahrenheit())
	fmt.Println(boilingC.ToFahrenheit())
}

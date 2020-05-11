package main

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
    "time"
    "unicode/utf8"
)

func rot13(in rune) rune {
    if in >= 'A' && in <= 'Z' {
        return ((((in - 'A') + 13) % 26) + 'A')
    }
    if in >= 'a' && in <= 'z' {
        return ((((in - 'a') + 13) % 26) + 'a')
    }
    return in
}

//Fazer funcionar
func manualMap(fun func(rune) rune, st string) string {
    var response string
    for ii := 0; ii < utf8.RuneCountInString(st); ii++ {
        response += string(fun(rune(ii)))
    }
    return response
}

func main() {
    fmt.Println("Packages")
    fmt.Println("\nStrings")

    a := "str"
    fmt.Println(strings.Contains(a, "m"))
    fmt.Println(strings.Contains(a, "t"))

    fmt.Println(strings.EqualFold("Victor", "victor"))
    fmt.Println(strings.EqualFold("Victor", "vict0r"))

    name := "Victorabarros"
    splitedName := strings.Split(name, "a")
    fmt.Println(splitedName)
    for ii := 0; ii < len(splitedName); ii++ {
        fmt.Println(splitedName[ii])
    }

    completeName := "Victor Almeida Barros"
    fmt.Println(strings.TrimSpace("  victor barros "))
    fmt.Println(completeName)
    fmt.Println(strings.ToLower(completeName))
    fmt.Println(strings.ToUpper(completeName))
    fmt.Println(strings.ToTitle(completeName))

    s := "cThis is a test 123 😃"
    s2 := strings.Map(rot13, s)
    fmt.Println(s2)
    s3 := strings.Map(rot13, s2)
    fmt.Println(s3)

    t := "cThit it a tett 123 😃"
    t2 := manualMap(rot13, t)
    fmt.Println(t2)
    t3 := manualMap(rot13, t2)
    fmt.Println(t3)

    fmt.Println("\nUTF8")
    fmt.Println("1234😃")
    fmt.Println(len("1234😃"))
    fmt.Println(utf8.RuneCountInString("1234😃"))

    fmt.Println("\ntime")
    now := time.Now()
    fmt.Println(now)
    fmt.Println(now.UnixNano()) //EPOCH https://www.epochconverter.com/

    fmt.Println("\nmath")
    var xx float64
    xx = 0.531
    fmt.Println(math.Cos(xx))
    fmt.Println(math.Pi)
    fmt.Println(math.Exp(-4))
    fmt.Println(math.Pow(10, -4))

    fmt.Println("\nmath/rand")
    rand.Seed(now.UnixNano())
    fmt.Println(rand.Intn(100))
    fmt.Println(rand.Intn(100))
}

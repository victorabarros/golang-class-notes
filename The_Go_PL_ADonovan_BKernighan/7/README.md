# 7 Interfaces

## 7.1 Interfaces as Contracts

```go
type Printer interface {
    Print()
}

type book struct {
    Name string
}

func (b book) Print() { fmt.Println("Book\t", b.Name) }

type game struct {
    Name string
}

func (g game) Print() { fmt.Println("Game\t", g.Name) }

func main() {
    stock := []Printer{
        book{"Moby Dick"},
        game{"GTA"},
    }

    for _, ii := range stock {
        ii.Print()
    }
}
```

## 7.2 Interface Types

```go
type Closer interface {
    Close() error
}

type PrintCloser interface {
    // Implemeting interface in other interface
    Printer
    Closer
}
```

## 7.3 Interface Satisfaction

> A type satisfies an interface if it possesses all the methodsthe interface requires.

an expression may be assigned to an interface only if its type satisfies the interface.

```go
var w io.Writer

    w = os.Stdout           // OK: *os.File has Write method
    w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
    w = time.Second         // compile error: time.Duration lacks Write method
```

## 7.4 Parsing Flags

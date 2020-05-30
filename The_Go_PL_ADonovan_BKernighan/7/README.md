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

> Example in .4/main.go

## 7.5 Interface Values

> Conceptually, a value of an interface type, or **interface value**, has two components, a concrete type and a value of that type.
> These are called the interface’s dynamic type and dynamic value.

## 7.6 Sorting

> Nice example. How explore type and methods:
> You can explore commons types new extensions methods to your particular rules.

```go
type StringSlice []string

func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
```

## 7.7 http.Hanlder interface

> Beatiful example at .7/main.go

## 7.8 error interface

`fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"`

## 7.9 Example: Expression Evaluator

> Nice webapp at .9/main.go

## 7.10 Tye Assertions

> A type assertion is an operation applied to an interface value.
> Syntactically, it looks like x.(T), where x is an expression of an interface typeand T is a type, called the “asserted” type.

```go
var w io.Writer // interface
w = os.Stdout
f := w.(*os.File)      // success: f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

## 7.11 Discriminating Errors with Type Assertions

## 7.12 Querying Behaviors with Interface Type Assertions

```go
package fmt

func formatOneValue(x interface{}) string {
    if err, ok := x.(error); ok {
        return err.Error()
    }
    if str, ok := x.(Stringer); ok { // As interface pattern name we can propose that `type Stringer interface {String() string}`
        return str.String()
    }
    // ...all other types...}
```

# Chapter 3 - Basic Data Typer

https://learning.oreilly.com/library/view/the-go-programming/9780134190570/ebook_split_027.html

## 3.1 Integers

> There are four distinct sizes ofsigned integers—8, 16, 32, and 64 bits—represented by the types int8, int16, int32, and int64, and corresponding unsigned versions uint8, uint16, uint32, and uint64.
> The type rune is a synonym for int32 and conventionally indicates that a value is a Unicode code point.
> Similarly, the type byte is an synonym for uint8.
> The range ofvalues of an n-bit number is from `−2**(n−1)` to `2**(n−1)−1`.
> The behavior of `/` depends on whether its operands are integers,so 5.0/4.0 is 1.25, but 5/4 is 1 because integer division truncates the result toward zero.

## 3.2 Floating-Point Numbers

## 3.3 Complex Numbers

## 3.4 Booleans

## 3.5 String

## 3.6 Constants

### 3.6.1 The Constant Generator `iota`

> A const declaration may use the _constant generator_ **iota**, which is used to create a sequence of related values without spelling out each one explicitly.
> An example from the time package, which defines named constants of type Weekday.
> Types of this kind are often called **enumerations** or **enums**.

```go
type weekday int

const (
    sunday    weekday = iota // 0
    monday                   // 1
    tuesday                  // 2
    wednesday                // 3
    thursday                 // 4
    friday                   // 5
    saturday                 // 6
)
```

Or more complex expressions:

```go
type enu int

const (
    en1  enu = 1 << iota // 1
    en2                  // 2
    en4                  // 4
    en8                  // 8
    en16                 // 16
)
```

> The `iota` mechanism has its limits. It’s not possible to generate the more familiar powers of 1000 (KB, MB, and so on) because there is no exponentiation operator.

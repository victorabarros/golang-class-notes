# 5 Functions

```go
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
```

## 5.3 Multiple Returns

> Bare Return:

```go
func xpto() (a int, err error) {
    a = 0
    err = fmt.Errorf("", nil)
    return // bare return
}
```

### 5.4.2 End of File EOF

```go
    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err == io.EOF {
            break // finished reading
        }
        if err != nil {
            return fmt.Errorf("read failed: %v", err)
        }
        // ...use r...
    }
```

## 5.9 Panic

> A panicis often the best thing to do when some “impossible” situationhappens.

```go
    switch s := suit(drawCard()); s {
        case "Spades":   // ...
        case "Hearts":   // ...
        case "Diamonds": // ...
        case "Clubs":    // ...
        default:
            panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
    }
```

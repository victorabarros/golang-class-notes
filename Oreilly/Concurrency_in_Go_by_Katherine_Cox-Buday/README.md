# Concurrency in Go
##### https://learning.oreilly.com/library/view/concurrency-in-go/9781491941294/
## Chapter 1
 - Moore’s Law
 > "is the observation that the number of transistors in a dense integrated circuit doubles about every two years." - https://en.wikipedia.org/wiki/Moore%27s_law
 - Race Condition
 > "are one of the most insidious types of concurrency bugs because they may not show up until years after the code has been placed into production" - https://learning.oreilly.com/a/concurrency-in-go/43600044/

[Example1](./main.go):
```go
    var data int
    go func() {
        data++ // line 3
    }()
    if data == 0 { // line 5
        fmt.Printf("the value is %v.\n", data) // line 6
    } else {
        fmt.Print("Data not Zero.\n")
    }
```
There is thre differences possibles answers:
```sh
    the value is 0.    // if line 3 runs after line 6
    the value is 1.    // if line 3 runs before line 6 but after line 5
    Data not Zero      // if line 3 runs before line 5
```

 - Atomicity
 > "When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible." - https://learning.oreilly.com/a/concurrency-in-go/43600097/

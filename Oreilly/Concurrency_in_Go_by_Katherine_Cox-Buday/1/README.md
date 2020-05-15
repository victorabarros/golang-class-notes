# Concurrency in Go
##### https://learning.oreilly.com/library/view/concurrency-in-go/9781491941294/
## Chapter 1
#### Moore’s Law
 > "is the observation that the number of transistors in a dense integrated circuit doubles about every two years." - https://en.wikipedia.org/wiki/Moore%27s_law
#### Race Condition
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

#### Atomicity
 > "When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible." - https://learning.oreilly.com/a/concurrency-in-go/43600097/

Example:
```go
    i++
```
It may look atomic, but a brief analysis reveals several operations:
 - Retrieve the value of i.
 - Increment the value of i.
 - Store the value of i.

While each of these operations alone is atomic, the combination of the three may not be, depending on your context.

> So why do we care? Atomicity is important because if something is atomic, implicitly it is safe within concurrent contexts. This allows us to compose logically correct programs, and—as we’ll later see—can even serve as a way to optimize concurrent programs.

#### Memory Access Synchronization
Example:
```go
    var data int
    go func() { data++}()
    if data == 0 {
        fmt.Println("the value is 0.")
    } else {
        fmt.Printf("the value is %v.\n", data)}
```
Remember that as it is written, there is a data race and the output of the program will be completely nondeterministic. In fact, there’s a name for a section of your program that needs exclusive access to a shared resource. This is called a **critical section**. In this example, we have three critical sections:
- Our goroutine, which is incrementing the data variables.
- Our if statement, which checks whether the value of data is 0.
- Our fmt.Printf statement, which retrieves the value of data for output.

> How solve this problem?
> Answer: **MUTEX**! Add `var mutex sync.Mutex` and `Lock/Unlock` each time `data` variable is called.

The calls to `Lock` can make our program slow. Every time we perform one of these operations, our program pauses for a period of time.
This brings up two questions:
 - _Are my critical sections entered and exited repeatedly?_
 - _What size should my critical sections be?_

Answering these two questions in the context of your program is an art, and this adds to the difficulty in synchronizing access to the memory.Synchronizing access to the memory also shares some problems with other techniques of modeling concurrent problems. We’ll discuss those in the next section.

#### Deadlocks, Livelocks, and Starvation

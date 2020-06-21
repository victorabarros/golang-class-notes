# 8 Goroutines and Channels

## 8.1 Goroutines

> The differencesbetween threads and goroutines are essentially quantitative, not qualitative.

## 8.2 Example: Concurrent Clock Server

## 8.4 Channels

### 8.4.1 Unbuffered Channels

> **Unbuffered channels** are sometimes called synchronous channels.
> When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.

### 8.4.2 Pipelines

```go
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; ; x++ { naturals <- x }
    }()

    // Squarer
    go func() {
        for {
            x := <-naturals
            squares <- x * x
        }
    }()

    // Printer (main goroutine)
    for { fmt.Println(<-squares) }
```

### 8.4.3 Unidirectional Channel Types

```go
func counter(out chan<- int) {
    for x := 0; x < 100; x++ { out <- x }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    // Atention to params types. It's unidirectional channels.
    // out is where the func send values: chan <- int
    // in  is where the func read values: <- chan int
    for v := range in { out <- v * v }
    close(out)
}

func printer(in <-chan int) {
    for v := range in { fmt.Println(v) }
}

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    go counter(naturals)
    go squarer(squares, naturals)
    printer(squares)
}
```

### 8.4.4 Buffered Channels

It's the channel that we assing the capacity on **make** built in method.

> A buffered channel has a queue of elements.
> The queue’s maximum size is determined when it is created, by the capacity argument to **make**.
> A send operation on a buffered channel inserts an element at the back of the queue, and a receive operation removes an element from the front.
> If the channel is full, the send operation blocks its goroutine until space is made available by another goroutine’s receive.
> Conversely, if the channel is empty, a receive operation blocks untila value is sent by another goroutine.

```go
func main() {
    ch = make(chan string, 3)        // {nil, nil, nil}
    ch <- "A"                        // {"A", nil, nil}
    ch <- "B"                        // {"A", "B", nil}
    ch <- "C"                        // {"A", "B", "C"}
    fmt.Println(<-ch)         // "A" // {"B", "C", nil}
    fmt.Println(cap(ch))      // "3"
    fmt.Println(len(ch))      // "2"
}
```

> Novices are sometimes tempted to use buffered channels within a single goroutine as a queue, lured bytheir pleasingly simple syntax, but this is a mistake.
> Channels are deeply connected to goroutine scheduling, and without another goroutine receiving from the channel, a sender—and perhaps the whole program—risks becoming blocked forever.
> If all you need is a simple queue, make one using a slice.

## 8.5 Looping in Parallel

> The correct way to optimize a bunk of jobs:

Before:

```go
func execubeChunkJobs(items []item) (resps []response) {
    for _ , ii := range items {
        resp, err := executeJob(ii) // Synchronous
        if err != nil {
            continue
        }
        resps = append(resps, resp)
    }
}
```

After:

```go
func execubeChunkJobsAssynchronous(items []item) (resps []response) {
    chResp := make(chan response)
    wg := sync.WaitGroup

    for _ , ii := range items {
        wg.Add(1)

        // Worker
        go func(ii item) {
            defer wg.Done()

            resp, err := executeJob(ii)
            if err != nil {
                return
            }
            chResp <- resp
        }(ii)
    }

    // Closer
    go func() {
        wg.Wait()
        close(chResp)
        for resp := range chResp {
            resps = append(resps, resp)
        }
    }
}
```

TODO: Run a server that only sleep(2 seconds), try benchmark with the executeChunk requesting this server.

## 8.6 Example: Concurrent Web Crawler

## 8.7 Multiplexing with select

## 8.8 Example: Concurrent Directory Traversal

## 8.10 Example: Chat Server

> Chat broadcast system.

## Final considerations

> It's a chapter that deserves a revisit later. Select, capacity, etc.. Are powerfulls and not trivials tools.

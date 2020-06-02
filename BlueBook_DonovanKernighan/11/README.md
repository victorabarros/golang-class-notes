# Chapter 11 - Tests

## 11.1

## 11.2

### 11.2.3 White-box testing

> "a white-box testhas privileged access to the internal functions and data structures ofthe package and can make observations and changes that an ordinaryclient cannot."

### 11.2.4 External test packages

### 11.2.5 Writing Effective tests

## 11.3 Coverage

## 11.4 Benchmark

`go test -bench=. -benchmem`
> The -benchmem command-line flagwill include memory allocation statistics in its report.
> Benchmarks like this tell us the absolute time required for a givenoperation, but in many settings the interesting performance questionsare about the relative timings of two different operations.For example, if a function takes 1ms to process 1,000 elements, howlong will it take to process 10,000 or a million?Such comparisons reveal the asymptotic growth of the running time ofthe function.Another example: what is the best size for an I/O buffer?Benchmarks of application throughput over a range of sizes can help uschoose the smallest buffer that delivers satisfactory performance.A third example: which algorithm performs best for a given job?Benchmarks that evaluate two different algorithms on the same inputdata can often show the strengths and weaknesses of each one on important orrepresentative workloads.
> Comparative benchmarks are just regular code.  They typically take the form of a singleparameterized function, called from several Benchmark functionswith different values, like this:

```go
func benchmark(b *testing.B, size int) { /* ... */ }
func Benchmark10(b *testing.B)   { benchmark(b, 10) }
func Benchmark100(b *testing.B)  { benchmark(b, 100) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
```

## 11.5 Profiling

> "We should forget about small efficiencies, say about 97% of thetime: premature optimization is the root of all evil." ~ Donald Knuth

- A **CPU profile** identifies the functions whose execution requiresthe most CPU time.The currently running thread on each CPU is interrupted periodicallyby the operating system every few milliseconds, with each interruptionrecording one profile event before normal execution resumes.
- A **heap profile** identifies the statements responsible forallocating the most memory.The profiling library samples calls to the internal memory allocation routines sothat on average, one profile event is recorded per 512KB of allocatedmemory.
- A **blocking profile** identifies the operations responsible forblocking goroutines the longest, such as system calls, channelsends and receives, and acquisitions of locks.The profiling library records an event every time a goroutine isblocked by one of these operations.

```sh
$ go test -run=NONE -bench=ClientServerParallelTLS64 \
    -cpuprofile=cpu.log net/http
PASS
BenchmarkClientServerParallelTLS64-8  1000
    3141325 ns/op  143010 B/op  1747 allocs/op
ok      net/http       3.395s


$ go tool pprof -text -nodecount=10 ./http.test cpu.log
2570ms of 3590ms total (71.59%)
Dropped 129 nodes (cum <= 17.95ms)
Showing top 10 nodes out of 166 (cum >= 60ms)
    flat    flat%   sum%    cum     cum%
    1730ms  48.19%  48.19%  1750ms  48.75%  crypto/elliptic.p256ReduceDegre
    230ms   6.41%   54.60%  250ms   6.96%   crypto/elliptic.p256Dif
    120ms   3.34%   57.94%  120ms   3.34%   math/big.addMulVV
    110ms   3.06%   61.00%  110ms   3.06%   syscall.Syscall
    90ms    2.51%   63.51%  1130ms  31.48%  crypto/elliptic.p256Square
    70ms    1.95%   65.46%  120ms   3.34%   runtime.scanobject
    60ms    1.67%   67.13%  830ms   23.12%  crypto/elliptic.p256Mul
    60ms    1.67%   68.80%  190ms   5.29%   math/big.nat.montgomery
    50ms    1.39%   70.19%  50ms    1.39%   crypto/elliptic.p256ReduceCarry
    50ms    1.39%   71.59%  60ms    1.67%   crypto/elliptic.p256Sum
```

# GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming

https://youtu.be/PTE4VJIdHPg?list=WL

- Always provide context
- Always provide justification

## Structuring your code and repository

Industrial serves highly mutable requiriments, then **Don't opt-in to structure before you need it**.

Some good ideias:

- Usually, a `cmd/` subdir for package mains
- Sometimes, a `pkg/` for other Go code
- Mode the business domain, not implementation
- But plenty for exceptions

## Program configuration

**Flags are the best way to configure your program.** Its self doncument the configuration surface area at runtime.

**Envs are useful in addition to flags.** Avoid using as the only way to configure.

## Component graph

### Prefer inversion-of-control dependenxy injection

Avoid dependency injection framworks or caniners. Ex:

**Bad =(**

```go
func BuildContainer() *dig.Container {
    container = dig.New()
    container.Provide(NewConfig)
    container.Provide(ConnectDatabase)
    container.Provide(NewPersonRepository)
    container.Provide(NewPersonService)
    container.Provide(NewServer)
    return container
}

func main() {
    container := BuildContainer()
    if err := container.Invoke(func(server *Server) {
        server.Run()
    }); err != nil {
        panic(err)
    }
}
```

**Good =D**

```go
func main() {
    cfg := GetConfig()
    db, err := ConnectDatabase(cfg.URN)
    if err != nil {
        panic(err)
    }
    repo = NewPersonRepository(db)
    service := NewPersonService(cfg.AccessToken, repo)
    server := NewServer(cfg.ListenAddr, service)
    server.Run()
}
```

## Goroutine lifecycle management

**Futures**

```go
func main() {
    future := make(chan int, 1)
    go func() { future <- process() }()
    result := <-future
}
```

**Async/await**

```go
func main() {
    c := make(chan int, 1)
    go func() { c <- process() }() // async
    v := <-c                       // await
}
```

**Scatter/gather**

```go
func main() {
    // Scatter
    c := make(chan result, 10)
    for i := 0; i <  cap(c); i++ {
        go func() {
            val, err := process
            c <- result{val, err}
        }
    }

    // Gather
    var total int
    for i := 0; i <  cap(c); i++ {
        res := <-c
        if res.err != nil {
            total += res.val
        }
    }
}
```

> A good programmer has strong command of these idioms.
> **A great programmer is proactive in teaching them to others**.

## Observability

> Better to
> **optimize for fast recovery (MTTR)**
> than for avoiding failure (MTBF)

> Metrics, logging and tracing are emergent
> **patterns of consumption**
> of observability data

### Metrics

- Counter (Add), Gauge (Set), Histogram (Observe)
- Most metrics systems have good Go client libs
- Prometheus is best-in-class
- Avoid host/check-based systems (Nagios, Ganglia)

### Logging

- Structured (JSON object) vs Unstructured (strings)
- Good looging libraries...
  - Use a small logger intergace, not a concrete object
  - Avoid package global state (global logger)
  - Enforce structured logging at calsites

### Tracing

- Standard is currently OpenTracinf (client API)
- Implementd by Zipkin, Jaeger, Datadog, etc...
- New approach: OpenCensus
- Relatively difficult to implement correctly

Metrics + logging + tracing
**-> All part of a complete observability breakfast**

## Testing

### A sanity check

> `go test ./...` should always succeed. If it has environment dependecies they must be opt-in test flag.
> Folowing a good pattern:

```go
func TestFunc(t *testing.T) {
    urn := os.Getenv("TEST_DB_URN")
    if urn == "" {
        t.Skip("set TEST_DB_URN to run this test")
    }

    db, _ := connect(urn)
    // ...
}
```

### Keep it simple

- Good observability is ultimately more importante
- Mitchell Hashimoto's [Advanced Testing with Go](https://youtu.be/8hQG7QlcLBk)
- Testing Pyramid: 80% unit tests, 20% integration
- Good unit tests are table-driven, use mocks
- Avoid introducing BDD/TDD "helper" packages until you experience a concrete need

## How much interface do I need?

> Go's type system is sctructural, not nominal:
> **interfaces are behavior contracts**
> thath generally belong in consuming code.
> **-> Accept interfaces**
> **and return structs**

Define interfaces at the major "fault lines" in your architecture:

- **Between func main and the rest**
- **Aling package API boundaries**

> Interfaces reify abstraction boudaries;
> **use them to help structure**
> **yout thourghts and design**
> especially when writing tests

## Context use misuse

> Using context for lifecycle management is
> **generally a good idea**
> and infectious

```go
// reportStore is a thin domain abstraction over GCS (Google Coud Storage)
type reportStore interface {
    listTimes(ctx context.Context, ...) (..., error)
    writeFile(ctx context.Context, ...) error
    serveFile(ctx context.Context, ...) error
}
```

> Using contexts for value propagation is
> **tricker and more dubious-**
> request-scoped data only, please

## Summary

> **Go has changed remarkably little**
> and I really appreciate that

> These were some of my experiences, but
> **I'm excited to hear yours, too!**
> Confereces are great for build empathy

@peterbourgon GopherConEU Jun 2018

# 9. Concurrency with Shared Variables

> We’ll point out some of the problems associated with sharing variables among multiple goroutines,
> the analytical techniques for recognizing those problems, and the patterns for solving them.
> Finally, we’ll explain some of the technical differences betweengoroutines and operating system threads.

## 9.1 Race Conditions

> Go mantra “Do not communicate by sharing memory; instead, share memory by communicating.”

TODO: Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program.
The result should indicate whether the transaction succeeded or failed due to insufficient funds.
The message sent to the monitor goroutine must contain both the amount to withdraw
and a new channel over which the monitor goroutine can send the boolean result back to Withdraw.

## 9.2 Mutual Exclusion: sync.Mutex

> Example of mutex using binary semaphore.

```go
var (
    semaphore = make(chan struct{}, 1)
    balance int
)

func Deposit(amount int) {
    semaphore <- struct{}{} // acquire token
    balance = balance + amount
    <-semaphore // release token
}

func Balance() int {
    semaphore <- struct{}{} // acquire token
    b := balance
    <-semaphore // release token
    return b
}
```

This pattern of mutual exclusion is so useful that it is supported directly by the **Mutex** type from the **sync package**.
Its **Lock** method acquires the token (called a lock) and its **Unlock** method releases it.
Each time a goroutine accesses the variables of the bank (just balance here), it must call the mutex’s Lock method to acquire an exclusive lock.

IPC: The mutex guards the shared variables. By convention, the variables guarded by a mutex are declared immediately after the declaration of the mutex itself. If you deviate from this, be sure to document it.

## 9.3 Read/Write Mutexes: sync.RWMutex

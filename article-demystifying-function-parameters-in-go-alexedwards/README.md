# demystifying function parameters in go

https://www.alexedwards.net/blog/demystifying-function-parameters-in-go

## notes

- difference between `paramenters` and `arguments`
- "in Go, **functions always operate on a copy of the arguments**. There are no exceptions to this."
- the paramenter `func example(param *int)` **param** is a **copy of the pointer**
- **reference** operator **&param** to get a **pointer** to the variable (remember, a pointer just contains a memory address)
- **dereference** operator ***param** to 'read through' and get the **underlying value** at that memory address
- "Go will automatically dereference the pointer for you when you use the dot operator . on it to access a field or call a method."

    ```go
    func incrementScore(p *player) {
        p.score += 10 // Go automatically does (*p).score += 10
    }
    ```

- "[Go does not support pass-by-reference behavior](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go). I've probably banged this drum enough already now, but parameters are always a copy of the arguments." [...] "this was removed over a decade ago (with the commit message Go has no '[reference types](https://codereview.appspot.com/8288044)')."
- "Parameters always contain a copy of the argument. Go doesn't have "reference types" or support pass-by-reference semantics."

## Summary (copy from article)

- Parameters always contain a copy of the argument. Go doesn't have "reference types" or support pass-by-reference semantics.
- For the basic Go types, as well as structs, arrays and functions, changing the value of a parameter in the function body won’t change the value of the argument. But if you do want to mutate the argument, you can use a pointer parameter instead and dereference it inside the function to ‘write-through’ a new value to the argument's memory address. For common operations on structs and arrays, Go will automatically dereference the pointer for you.
- Because of the way that they're implemented by the Go runtime, changes you make to map, slice, channel parameters in a function will mutate the argument. If you don't want this, make a clone at the start of the function and use that instead.
- Using the = operator to assign a new value to a parameter does not affect the argument (unless you are manually-or-automatically dereferencing a pointer and 'writing-through' a new value). So for slices, if you want a function to perform an append or reslice operation that mutates the argument, you should use a pointer to a slice as the function parameter and dereference it as necessary.

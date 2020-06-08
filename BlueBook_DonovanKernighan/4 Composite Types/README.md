# Chapter 4 - Composite Types

> Arrays and structs are aggregate types; their values are concatenations of other values in memory.
> Arrays are homogeneous — their elements all have the same type — whereas structs are heterogeneous.
> Both arrays and structs are fixed size.In contrast, slices and maps are dynamic data structures that grow asvalues are added.

## 4.1 Array

“...” appears in place ofthe length, the array length is determined by the number of initializers.

```go
    q := [...]int{1, 2, 3}
    fmt.Printf("%T\n", q) // "[3]int"
    q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
```

```go
    months := [...]string{1: "January", /* ... */, 12: "December"}
```

January is months[1] and December is months[12]. Ordinarily, the array element at index 0 would contain "" (ampty string).

## 4.2 Slices

> A slice has three components: a pointer, a length, and acapacity.

```go
// reverse reverses a slice of ints in place.
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

    // Here we reverse the whole array a:
    a := [...]int{0, 1, 2, 3, 4, 5}
    reverse(a[:])
    fmt.Println(a) // "[5 4 3 2 1 0]"
```

> Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contain the same elements.

```go
    var s []int    // len(s) == 0, s == nil
    s = nil        // len(s) == 0, s == nil
    s = []int(nil) // len(s) == 0, s == nil
    s = []int{}    // len(s) == 0, s != nil
```

## 4.3 Maps

> In Go, a map is a referenceto a hash table.

A map element is not a variable, and we cannot take its address:
> One reason that we can’t take the address of a map element is that growing a map might cause rehashing of existing elements into new storage locations,thus potentially invalidating the address.

```go
    _ = &ages["bob"] // compile error: cannot take address of map element

```

> This is a common pattern:

```go
import "sort"

func main() {
    ages := map[string]int{
        "charlie": 34,
        "alice":   31,
    }

    var names [len(ages)]string
    for name := range ages { // maps are unordereds. Every iteration will be in random order.
        names = append(names, name)
    }

    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
}
```

```go
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }

    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv { // elegant
            return false
        }
    }
    return true
}
```

## 4.4 Struct


# 10. Packages and the Go Tool

TODO: How is `func init()` behaviour? Does it inicialized befores `func main()`?

## 10.6 Packages and Naming

> how to follow Go’s distinctive conventions for naming packages and their members.

```go
package strings

type Replacer struct{ /* ... */ }
func NewReplacer(oldnew ...string) *Replacer

type Reader struct{ /* ... */ }
func NewReader(s string) *Reader
```

```go
package rand

type Rand struct{ /* ... */ }
func New(oldnew ...string) *Replacer
```

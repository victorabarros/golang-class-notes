# 6 Methods

## 6.1 Declarations

```go
    type Point struct {X, Y float64}

    // method of the Point type
    func (p Point) Distance(q Point) float64 {
        return math.Hypot(q.X-p.X, q.Y-p.Y)
    }

    func main() {
        fmt.Println(Distance(p, q)) // function
        fmt.Println(p.Distance(q))  // method
    }

    // traditional function
    func Distance(p, q Point) float64 {
        return math.Hypot(q.X-p.X, q.Y-p.Y)
    }
```

## 6.2 Methods with Pointer Receiver

> if a function needs to update a variable, or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable using a pointer.

```go
    func (p *Point) ScaleBy(factor float64) {
        p.X *= factor
        p.Y *= factor
    }
```

## 6.3 Composing Types by Struct Embedding

```go
import "image/color"

type ColoredPoint struct {
    Point
    Color color.RGBA
}

func main() {
    cp1 := ColoredPoint{}
    cp1.Point.Y = 2
    fmt.Println(cp1.Y) //Shortcut cp1.Point
    cp1.ScaleBy(2) //Shortcut to Point.ScaleBy

    cp2 := ColoredPoint{}
    cp1.Distance(cp2) // compile error: cannot use q (ColoredPoint) as PointA ColoredPoint
}
```

Readeble example of mutex:

```go
var (
    mu sync.Mutex // guards mapping
    mapping = make(map[string]string)
)

func Lookup(key string) string {
    mu.Lock()
    v := mapping[key]
    mu.Unlock()
    return v
}

// The version below groups together the two related variables in a single package-level variable, cache
// gives more expressive names to the variables relatedto the cache

var cache = struct {
    sync.Mutex
    mapping map[string]string
} {
    mapping: make(map[string]string),
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
```

## 6.4 Method Values and Expressions

```go
    p := Point{1, 2}
    q := Point{4, 6}
    distance := Point.Distance   // method expression
    fmt.Println(distance(p, q))  // "5"
    fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
```

## 6.5 Bit Vector Type

> Bit vector is other alternative to "set" structure, besides map[T]bool.
> But is much more complex.

## 6.6 Encapsulation

> When naming a getter method, we usually omit the Get prefix.
> This preference for brevity extends to all methods, not just field accessors, and to other redundant prefixes as well, such as Fetch, Find, and Lookup


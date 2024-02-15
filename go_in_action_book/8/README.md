# Chapter 8

## Methods

Difference between `func (u user)` and `func (u *user)` behavior.

```go
package main

import "fmt"

type user struct {
	name string
}

func newUser(name string) user {
	return user{
		name: name,
	}
}

func (u user) changeName(newName string) {
	u.name = newName
}

func (u *user) changeNamev2(newName string) {
	u.name = newName
}

func main() {
	vic := newUser("victor")

	fmt.Printf("%+2v\n", vic)

	vic.changeName("vanessa")
	fmt.Printf("%+2v\n", vic)

	vic.changeNamev2("vanessa")
	fmt.Printf("%+2v\n", vic)
}
```

out:

```log
{name:victor}
{name:victor}
{name:vanessa}
```

### if use map instead of struct

```go
package main

import "fmt"

type user map[string]string

func newUser(name string) user {
	return user{
		"name": name,
	}
}

func (u user) changeName(newName string) {
	u["name"] = newName
}

func (u *user) changeNamev2(newName string) {
	(*u)["name"] = newName
}

func main() {
	vic := newUser("victor")

	fmt.Printf("%+2v\n", vic)

	vic.changeName("vanessa")
	fmt.Printf("%+2v\n", vic)

	vic.changeNamev2("vanessa")
	fmt.Printf("%+2v\n", vic)
}
```

out:

```log
map[name:victor]
map[name:vanessa]
map[name:vanessa]
```

### if use array instead of struct

```go
package main

import "fmt"

type user [1]string

func newUser(name string) user {
	return user{
		name,
	}
}

func (u user) changeName(newName string) {
	u[0] = newName
}

func (u *user) changeNamev2(newName string) {
	u[0] = newName
}

func main() {
	vic := newUser("victor")

	fmt.Printf("%+2v\n", vic)

	vic.changeName("vanessa")
	fmt.Printf("%+2v\n", vic)

	vic.changeNamev2("vanessa")
	fmt.Printf("%+2v\n", vic)
}
```

out:

```log
[victor]
[victor]
[vanessa]
```

### if use slice instead of struct

```go
package main

import "fmt"

type user []string

func newUser(name string) user {
	return user{
		name,
	}
}

func (u user) changeName(newName string) {
	u[0] = newName
}

func (u *user) changeNamev2(newName string) {
	(*u)[0] = newName
}

func main() {
	vic := newUser("victor")

	fmt.Printf("%+2v\n", vic)

	vic.changeName("vanessa")
	fmt.Printf("%+2v\n", vic)

	vic.changeNamev2("vanessa")
	fmt.Printf("%+2v\n", vic)
}
```

out:

```log
[victor]
[vanessa]
[vanessa]
```

### conclusion

|type|output|behaviour|
-|-|-
struct|{name:victor}<br/>{name:victor}<br/>{name:vanessa}|struct use copy as default
map<br/>**map[string]string**|map[name:victor]<br/>map[name:vanessa]<br/>map[name:vanessa]|map use reference as default
array<br/>**[1]string**|[victor]<br/>[victor]<br/>[vanessa]|freaking weird. don't know the answer
slice<br/>**[]string**|[victor]<br/>[vanessa]<br/>[vanessa]|map use reference as default

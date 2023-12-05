package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func New(value int) tree {
	return tree{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (t *tree) Add(newTree tree) {
	fmt.Println(1, t)
	if t == nil {
		t = &newTree
		fmt.Println(7, t)
		return
	}
	fmt.Println(2, t)

	if newTree.value < t.value {
		fmt.Println(3, t)
		t.left.Add(newTree)
	} else if newTree.value == t.value {
		fmt.Println(4, t)
		// do nothing
	} else {
		fmt.Println(5, t)
		t.right.Add(newTree)
	}
	fmt.Println(6, t)
}

func main() {
	newTree := New(5)
	// for i := 0; i < 10; i++ {
	// 	newTree.Add(New(i))
	// }
	newTree.Add(New(4))

	fmt.Printf("%+2v\n", newTree)
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	AppendValues(values[:0], root)
}

// AppendValues appends the elements of t to values in order
// and returns the resulting slice.
func AppendValues(values []int, t *tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t

	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

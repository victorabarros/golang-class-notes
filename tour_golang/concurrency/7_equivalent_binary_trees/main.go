// https://tour.golang.org/concurrency/7
package main

// Tree struct { Left  *Tree; Value int; Right *Tree}
import (
    "fmt"
    "golang.org/x/tour/tree"
    "time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    walkTree(t, ch)
    // close(ch)
}
func walkTree(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        walkTree(t.Left, ch)
    }
    // Add valeu to cha in the middle of the change left to right.
    // Ensuring sort to channel values.
    ch <- t.Value
    if t.Right != nil {
        walkTree(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func printChan(ch chan int) {
    for i := range ch {
        fmt.Println(i)
    }
}

func main() {
    t := tree.New(1)
    ch := make(chan int)
    go Walk(t, ch)
    go printChan(ch)
    go printChan(ch)
    time.Sleep(99999999)
}

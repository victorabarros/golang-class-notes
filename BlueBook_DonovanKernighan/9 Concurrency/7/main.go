package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex //guards cache
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a new Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// GetThreadNotSafe makes a internal get from Memo
func (memo *Memo) GetThreadNotSafe(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

// GetWithTreadSafeButInneficient makes a internal get from Memo
func (memo *Memo) GetWithTreadSafeButInneficient(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

// GetWithThreadSafeAndEfficient makes a internal get from Memo
func (memo *Memo) GetWithThreadSafeAndEfficient(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A MemoWithEntry caches the results of calling a Func.
type MemoWithEntry struct {
	f     Func
	mu    sync.Mutex //guards cache
	cache map[string]*entry
}

// NewMemoWithEntry returns a new Memo
func NewMemoWithEntry(f Func) *MemoWithEntry {
	return &MemoWithEntry{f: f, cache: make(map[string]*entry)}
}

// Get ...
func (memo *MemoWithEntry) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() map[string]bool {
	return map[string]bool{
		"http://ota-api.hud":  true,
		"http://ota-api.hud/": true,
	}
}

func main() {
	start := time.Now()

	// fmt.Println("\nEXAMPLE1")
	// start = time.Now()
	// example1()
	// fmt.Printf("%s, %s\n", "EXAMPLE1", time.Since(start))

	// fmt.Println("\nEXAMPLE2")
	// start = time.Now()
	// example2()
	// fmt.Printf("%s, %s\n", "EXAMPLE2", time.Since(start))

	fmt.Println("\nEXAMPLE3")
	start = time.Now()
	example3()
	fmt.Printf("%s, %s\n", "EXAMPLE3", time.Since(start))

	fmt.Println("\nEXAMPLE4")
	start = time.Now()
	example4()
	fmt.Printf("%s, %s\n", "EXAMPLE4", time.Since(start))
}

func example1() {
	m := New(httpGetBody)

	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.GetThreadNotSafe(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func example2() {
	m := New(httpGetBody)
	var n sync.WaitGroup

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.GetWithTreadSafeButInneficient(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func example3() {
	m := New(httpGetBody)
	var n sync.WaitGroup

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.GetWithThreadSafeAndEfficient(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func example4() {
	m := NewMemoWithEntry(httpGetBody)
	var n sync.WaitGroup

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

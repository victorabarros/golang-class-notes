package main

import (
	"fmt"
	"net/http"
)

type fetcher interface {
	fetch() string
}

func main() {}

func server(f fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, f.fetch())
	}
}

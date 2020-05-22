// Learn Go in 3 hours
// Section 2, Video 1
// Our First Go Program
package main

import (
	"fmt"
	"net/http"
)

/*
   All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	// GET /hello
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("/hello")
		name := req.URL.Query().Get("name")
		// how validate header?
		rw.Write([]byte(fmt.Sprintf("Hello, %s \n", name)))
	})

	// GET /print
	http.HandleFunc("/print", func(rw http.ResponseWriter,
		req *http.Request) {
		fmt.Println("/print")
		info := req.URL.Query().Get("info")
		fmt.Println(req.Header)
		fmt.Println(info, rw.Header())
	})

	http.ListenAndServe(":8081", nil)
}

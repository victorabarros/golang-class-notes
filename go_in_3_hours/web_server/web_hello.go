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
        name := req.URL.Query().Get("name")
        rw.Write([]byte(fmt.Sprintf("Hello, %s \n", name)))
    })

    // GET /print
    http.HandleFunc("/print", func(rw http.ResponseWriter, req *http.Request) {
        info := req.URL.Query().Get("info")
        rw.Write([]byte(fmt.Sprintf("%s \n", info)))
    })

    http.ListenAndServe(":8080", nil)
}

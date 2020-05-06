package main

import (
    "fmt"
	// "bytes"
	"io/ioutil"
    "net/http"
    // "time"
)

func main() {
    // buf := bytes.NewReader([]byte(`{"msg": 1234}`))
    // How add header?
    for {
        go getWithHeader()
    }
    return
}

func getWithHeader() {
    client := &http.Client{}

    req, err := http.NewRequest(
        "GET", "http://localhost:8081/print?info=hotel", nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Add("If-None-Match", `W/"wyzzy"`)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        fmt.Println(err)
    }
}

func getDefault() {
    resp, err := http.Get("http://localhost:8081/print?info=hotel")
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}

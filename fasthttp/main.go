package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// var (
//     w http.ResponseWriter
//     r *http.Request
//     ctx *fasthttp.RequestCtx
// )

var (
	fastClient = fasthttp.Client{}
)

func main() {
	req := fasthttp.AcquireRequest()
	req.Header.Add("X-Header", "0")

	req.SetRequestURI("http://google.com/")

	resp := fasthttp.AcquireResponse()

	fastClient.Do(req, resp)

	fmt.Println(string(req.RequestURI()))
	fmt.Println(string(resp.Body()))
}

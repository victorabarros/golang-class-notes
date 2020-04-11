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

func main() {
    fmt.Println("Start")
    // statusCode, _, err := fasthttp.Get(nil, "http://google.com/")
    // fmt.Println(statusCode, err)

    // headers := fasthttp.RequestHeader{}
    // headers.Add("Ota-Cache-Refresh", "0")
	req := fasthttp.AcquireRequest()
	req.Header.Add("Ota-Cache-Refresh", "0")
	resp := fasthttp.AcquireResponse()

	endpoint := fmt.Sprintf("%shotels/%s-%s",
                            "http://www.ota-api.hud/",
                            "OMN",
                            "2020")

    req.SetRequestURI(endpoint)
	// ctx := fasthttp.RequestCtx{Request: req}
	client := fasthttp.Client{}

	fmt.Println(string(req.RequestURI()))
	client.Do(req, resp)
	fmt.Println(string(resp.Body()))
}

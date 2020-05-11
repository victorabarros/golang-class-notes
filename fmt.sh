docker run -v $(pwd):/go/src/github.com/test/ -w /go/src/github.com/test/ golang:1.14 gofmt -e -l -s -w .

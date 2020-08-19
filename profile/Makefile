all: clean build exec prof

build:
	@go build main.go

clean:
	@rm -rf main report/*

exec:
	@./main

prof:
	@go tool pprof -pdf $(GOPATH)/src/github.com/victorabarros/golang-class-notes/profile/main ./report/cpuprofile > ./report/cpuprofile.pdf

graph:
	@echo "make graph"
	@apt update
	@apt install -y graphviz

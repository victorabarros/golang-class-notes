all: clean build exec pdf

build:
	@go build main.go

clean:
	@rm -rf main report/*

exec:
	@./main

pdf:
	@go tool pprof -pdf ./main ./report/cpuprofile > ./report/cpuprofile.pdf
	@go tool pprof -pdf ./main ./report/memprofile > ./report/memprofile.pdf

prof-cpu:
	@go tool pprof ./main ./report/cpuprofile

prof-mem:
	@go tool pprof ./main ./report/memprofile

graph:
	@echo "make graph"
	@apt update
	@apt install -y graphviz

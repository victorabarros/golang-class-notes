rm -rf bench.log
go test -bench=. -benchmem > bench.log

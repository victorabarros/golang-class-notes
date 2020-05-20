rm -rf c.out report.log
go test -bench=. -benchmem > report.log

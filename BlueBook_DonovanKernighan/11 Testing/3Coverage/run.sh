rm -rf c.out
go test -cover -coverprofile=c.out
go tool cover -html=c.out

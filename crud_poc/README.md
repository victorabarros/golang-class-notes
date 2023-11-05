# POC to create a CRUD

## Quick start

```sh
# to start to project
go mod init github.com/victorabarros/Learning

# to download the api framework
go get github.com/gin-gonic/gin

# to create the main.go file
mkdir cmd ; mkdir cmd/api ; touch cmd/api/main.go

# to run the project
go run cmd/api/main.go

# to download the postgres driver
go get github.com/jackc/pgx/v4

# to create the connection file
mkdir internal ; mkdir internal/database ; touch internal/database/connection.go

# remove old version of pgx
go get github.com/jackc/pgx/v4@none

# add new version
go get github.com/jackc/pgx/v5
```

## references

https://youtu.be/9BeFJuzg_yw

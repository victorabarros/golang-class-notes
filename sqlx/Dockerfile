# TODO: alpine or slim, whos better?
FROM golang:1.14-alpine

WORKDIR src/github.com/vabarros/golang-sqlx/.
COPY . .

RUN apk add git && \
    go get -d -v github.com/jmoiron/sqlx && \
    go get -d -v github.com/go-sql-driver/mysql && \
    go build main.go

CMD ["./main"]

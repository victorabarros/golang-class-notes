FROM golang:1.14.2
    WORKDIR ${GOPATH}/src/github.com/victorabarros/go_mod/
    COPY . ./

    # Cria go.mod sem dependencias
    RUN go mod init

    # Varre o código, escreve as dependencias no go.mod e go.sum
    RUN go mod tidy

    # Depois compila
    RUN go build -o bin/main main.go

    CMD ["./bin/main"]

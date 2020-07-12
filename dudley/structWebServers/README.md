# How I Structure Web Servers in Go

> https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers

## Directory Structure

```
├── app/                    # entry point newcomers gravitate towards when exploring the codebase
|   └── service-api/        # micro-service API for this repository; all HTTP implementation details live here
|       ├── cfg/            # configuration files, usually json or yaml saved in plain text files, as they should be checked into git too
|       ├── middleware/     # for all middleware
|       ├── routes/         # API application’s RESTFul-like surface
|       |   ├── makes/
|       |   |   └── models/**
|       |   ├── create.go
|       |   ├── create_test.go
|       |   ├── get.go
|       |   └── get_test.go
|       ├── webserver/      # contains all shared HTTP structs and interfaces (Broker, configuration, Server, etc)
|       ├── main.go         # bootstrapped (New(), Start())
|       └── routebinds.go   # BindRoutes() function
├── cmd/                    # where any command-line applications belong
|   └── service-tool-x/
├── internal/               # directory that cannot be imported by projects outside of this repo
|   └── service/            # domain logic; it can be imported by service-api
|       └── mock/
└── pkg/                    # packages that are encouraged to be imported by projects outside this repo
    ├── client/             # library for accessing service-api. Other teams can import it without having to write their own
    └── dtos/               # data transfer objects, structs designed for sharing data between packages and encoding/transmitting. /internal/service is responsible for mapping the DTOs to/from its internal models
```

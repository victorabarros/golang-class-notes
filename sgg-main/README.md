# Sensys Gatso - Technical Assignment

## Project structure

The project sticks to the project layout described in https://github.com/golang-standards/project-layout

- `cmd/articles` - main file of the project.
- `internal/` - internal files of the project. Not intended to be shared or imported by other projects.
- `pkg/` - files that can eventually be shared by other projects. E.g., models, middlewares, and error codes.
- `scripts/` - helper scripts to automate some tasks.
- `vendor/` - external dependencies. Whether it's versioned or not depends on each team. I prefer to version the vendor folder and prevent headaches if someone deletes an external dependency.

## Solution

The project consists in an application where users can create articles with some predefined fields (title, description, expiration date, and images). The following tasks can be accomplished using HTTP requests and :

- Create an article (with or without images)
- Attach images to an existing article
- Check the size of the image before storing it
- List the existing articles also filtering by images
- Remove expired articles automaticaly

The core of the project can be described with the following layers: repository, service, and transport. The requests that are sent to the application go through the following path:

1. (incoming request) -> *transport* -> *service* -> *repository*
2. (store, retrieve or update the data)
3. *repository* -> *service* -> *transport* -> (response)

All layers are independent from each other in terms of implementation details. Since I'm using interfaces, we can replace an entire layer without touching the code of other layers. For example, we can replace the concrete implementation of a database (e.g., Postgres or MySQL) or add or remove transport protocols without changing one single line of the other layers.

### Repository layer

Repository is the inner most layer and has all the logic related to the data persistence. The part of the code related to the repository can be found in the package: `internal/data`. There is an interface that all concrete implementations of the repository must implement (e.g., in-memory, postgres, mysql, etc.):

```go
type ArticleRepository interface {
	Create(ctx context.Context, article model.Article) (*model.Article, error)
	Retrieve(ctx context.Context, articleID uuid.UUID) (*model.Article, error)
	Update(ctx context.Context, article model.Article) (*model.Article, error)
	List(ctx context.Context) ([]model.Article, error)
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, articleID uuid.UUID) error
}
```

The concrete implementation for this project is an in-memory database.

### Service layer

The service layer is where resides the business logic of the application. The code related to the service layer is found in: `internal/service`. There is also a high-level interface for the service layer:

```go
type ArticleService interface {
	CreateArticle(ctx context.Context, article model.Article) (*model.Article, error)
	AttachImage(ctx context.Context, articleID uuid.UUID, image model.Image) (*model.Article, error)
	ListArticles(ctx context.Context) ([]model.Article, error)
	RetrieveArticleByID(ctx context.Context, articleID uuid.UUID) (*model.Article, error)
	CheckExpired(ctx context.Context) (int, error)
	CheckImageSize(ctx context.Context, url string) error
}
```

### Transport layer

This is where resides the code responsible by the handlers for each transport (HTTP and gRPC). The code for this can be found in `internal/transport`. In this project there are two implementations: `http` package to keep it compatible with the popular and widely used communication protocol; and `grpc` for high-performance remote procedure calls between microservices.

### Other important features

The package `internal/server` contains the implementation of concrete servers where I keep independent implementations separate from each other. In that way, I keep a HTTP server, a gRPC server, and a cron server. Those servers are created and started in the package `cmd/articles`. This is where resides the code that will start the whole application.


The cron server runs a job that automatically deletes expired articles. The cron job will kick off every day at midnight and run a check to delete the articles that expired the previous day.

Another important feature is the validation of the image size. Before saving an image, a HEAD request is sent to the image to verify the content lengh. If it's bigger than 5MB, then the request fails.

### Service configuration

The service configuration uses environment variables. In that way, we can explicitly set environment variables to change the service configuration. That can be accomplished in different ways, depending on the environment the service is running (e.g., Kubernetes).

This is the struct that holds the environment variables configurations. If now variable is set in the environment, then the default values are taken:

```go
type Configuration struct {
	Prefix string `env:"PREFIX" env-default:"/api/v1"`
	Cron   struct {
		DeleteArticles string `env:"CRON_DELETE_JOBS" env-default:"0 0 0 * * *"` // kicks-off at midnight everyday
	}
	HTTP struct {
		Port           int           `env:"HTTP_PORT" env-default:"8080"`
		DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT" env-default:"30s"`
	}
	GRPC struct {
		Port int `env:"GRPC_PORT" env-default:"9090"`
	}
}
```

### Tests

The code contains unit and integration tests that can be run with the following command:

```bash
make test
```

## Makefile useful goals

- `build`: Build executales
- `clean`: Clean
- `fmt`: Execute go fmt
- `lint`: Run static checks
- `protoc`: Protobuf compile
- `test`: Run tests
- `vendor`: Fetch dependencies
- `reinstall-golangci-lint`: Remove any existing versions of golangci-lint and install it again
- `run`: Run executable
- `help`: Display help screen

## How to set up the environment

Requirements:
- Go version [go1.18+](https://go.dev/dl/)
- curl (for http requests)
- Protocol buffer compiler (protoc) version 3
- [grpcurl](https://github.com/fullstorydev/grpcurl) (for gRPC requests)

For `protoc`, there is a Makefile target. It works with Linux and OSX and you just need to run:
```bash
make protoc
```
The command above will download the protocol buffer compiler and it will be ready for use. The command also compiles the `*.proto` file in the repository and generates the `.go` files that are necessary for the project to build.

Now, you can build the project by running:
```bash
make build
```
That command will generate a binary named `articles-api` that you can execute to start the project. The binary can also be executed by using the `run` goal from the Makefile:
```bash
make run
```

The `grpcurl` program is used to execute the gRPC requests. The program can be installed with:
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
> **Note**
> It's important to note that your `GOPATH` needs to be set up and the go binaries need to be exported in your `PATH`.

## Executing the requests

### HTTP

> **Note**
> If you prefer, you can append ` | jq .` to the end of your request to see a nice JSON representation.

Retrieve a list of all articles:
```bash
curl -XGET -L "http://localhost:8080/api/v1/articles"
```

Retrieve a list of articles with images:
```bash
curl -XGET -L "http://localhost:8080/api/v1/articles?withImages=yes"
```

Retrieve a list of articles without images:
```bash
curl -XGET -L "http://localhost:8080/api/v1/articles?withImages=no"
```

Create article without images:
```bash
curl -XPOST -L "http://localhost:8080/api/v1/articles" -d '{"title":"a","description":"desc","expirationDate":"2022-10-23"}'
```

Create article with image (<5MB):
```bash
curl -XPOST -L "http://localhost:8080/api/v1/articles" -d '{"title":"a","description":"d","expirationDate":"2022-10-25","images":[{"path":"https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Snake_River_%285mb%29.jpg/512px-Snake_River_%285mb%29.jpg"}]}'
```

Create article with image (>5MB):
```bash
curl -XPOST -L "http://localhost:8080/api/v1/articles" -d '{"title":"a","description":"d","expirationDate":"2022-10-25","images":[{"path":"https://upload.wikimedia.org/wikipedia/commons/2/2d/Snake_River_%285mb%29.jpg"}]}'
```

Attach image (needs a valid article UUID):
```bash
curl -XPOST -L "http://localhost:8080/api/v1/articles/21e66719-1adf-4c87-9af6-b6c985b69c2a/images" -d '{"title":"a","description":"d","expirationDate":"2022-10-25","images":[{"path":"https://upload.wikimedia.org/wikipedia/commons/2/2d/Snake_River_%285mb%29.jpg"}]}'
```

### gRPC

List all articles:
```bash
grpcurl -plaintext  localhost:9090 Articles/List
```

List articles with images:
```bash
grpcurl -plaintext -d '{"withImages": "yes"}'  localhost:9090 Articles/List
```

List articles without images:
```bash
grpcurl -plaintext -d '{"withImages": "no"}'  localhost:9090 Articles/List
```

Create article without image:
```bash
grpcurl -plaintext -d '{"title":"a","description":"d","expirationDate":"2022-10-25"}' localhost:9090 Articles/Create
```

Create article with image (>5MB):
```bash
grpcurl -plaintext -d '{"title":"a","description":"d","expirationDate":"2022-10-25","images":[{"path":"https://upload.wikimedia.org/wikipedia/commons/2/2d/Snake_River_%285mb%29.jpg"}]}' localhost:9090 Articles/Create
```

Create article with image (<5MB):
```bash
grpcurl -plaintext -d '{"title":"a","description":"d","expirationDate":"2022-10-25","images":[{"path":"https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Snake_River_%285mb%29.jpg/512px-Snake_River_%285mb%29.jpg"}]}' localhost:9090 Articles/Create
```

Attach image:
```bash
grpcurl -plaintext -d '{"articleId":"e0beb281-c79e-4d36-be14-1b7618162e7e","path":"https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Snake_River_%285mb%29.jpg/512px-Snake_River_%285mb%29.jpg"}' localhost:9090 Articles/AttachImage
```

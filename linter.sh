# https://github.com/golangci/golangci-lint
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run -v

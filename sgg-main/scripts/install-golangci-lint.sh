#!/usr/bin/env bash

# Install golangci-lint
echo "==> Install golangci-lint"

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin $GOLANGCILINT_VERSION

exit 0

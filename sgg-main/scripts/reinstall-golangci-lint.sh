#!/usr/bin/env bash

# Remove and install golangci-lint
echo "==> Reinstall golangci-lint"

echo "Removing golangci-lint"
rm -f $GOLANGCILINT &>/dev/null
echo "Installing golangci-lint $GOLANGCILINT_VERSION"
make $GOLANGCILINT

exit 0

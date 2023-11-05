#!/usr/bin/env bash

# Run static check with golangci-lint
echo "==> Run golangci-lint"

if [[ -x "$(command -v $GOLANGCILINT)" ]]; then
  current_version=$($GOLANGCILINT --version | perl -lne "/\d+\.\d+\.\d+/ and print $&")
  if [[ "v$current_version" != $GOLANGCILINT_VERSION ]]; then
    printf "\033[0;31mWrong version of golangci-lint installed.\nPlease install version $GOLANGCILINT_VERSION\033[0m\n"
    exit 1
  fi
fi

$GOLANGCILINT run -v -c .golangci.yml

exit 0

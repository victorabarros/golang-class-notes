#!/usr/bin/env bash

# Installing protoc

echo "==> Installing protobuf compiler."

case "$OSTYPE" in
  darwin*)
    OS="osx"
    ARCH="universal_binary"
    ;;
  linux*)
    OS="linux"
    ARCH=$(arch)
    ;;
  *)
    printf "\033[0;31mUnknown os type: $OSTYPE\033[0m\n"
    echo "Expects linux or darwin (osx)"

    exit 1
  ;;
esac
PROTOC_RELEASE=protoc-$PROTOC_VERSION-$OS-$ARCH.zip
wget "https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_RELEASE"
unzip -o $PROTOC_RELEASE -d ./protoc
rm -f $PROTOC_RELEASE
chmod -R 775 ./protoc/bin

go install google.golang.org/protobuf/cmd/protoc-gen-go@$PROTOC_GEN_GO
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$PROTOC_GEN_GO_GRPC

exit 0

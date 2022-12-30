# CLI with Cobra

## start

use `make debug`

```sh
go mod init github.com/victorabarros/cli-with-cobra
go install github.com/spf13/cobra-cli@latest
cobra-cli init
```

## develop

add `fmt.Println("Print: " + strings.Join(args, " "))` to command run

## how to run

macos: `env GOOS=darwin GOARCH=amd64 go build`

## install cli

`go install`
<!-- > Note: make sure that `$GOPATH/bin` is on your `$PATH` environment variable, because is where it will be installed. -->
<!-- `export PATH=$GOPATH/bin:$PATH` -->

## usage

```log
$ cli-with-cobra "print this _|_"
Print: print this _|_
```

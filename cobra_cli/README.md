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

## install cli

go install

## usage

```log
$ cli-with-cobra "print this _|_"
Print: print this _|_
```

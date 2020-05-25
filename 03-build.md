# Building

## About

`go help build`

## GOPATH

`go build` runs in the context of the `GOPATH` which is defaulted at `${HOME}/go`.
Change this env var in the shell you're running the build to change the `go build`
command's behaviour.

## Force Output

`go build -o <target> <source>`

This will force writing libraries to disk (by default, `go build` doesn't do that).

## Intermediate Depenndecies

`go build -i <target>`

This will build intermediate dependencies. For instance, it will build imported packages
as libraries, besides building the final executable.

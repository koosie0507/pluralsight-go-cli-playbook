# Testing

## About

* `go help test`
* `go help testflag`
* name your files with the suffix `_test` and the compiler will leave them
  out of your prod build

## Running Tests

* run tests in one package: `go test mylib`
* run tests in package + subpackages: `go test mylib/...`
* pass arguments to the binary the tests will be testing: `go test -args`
* limit CPUs available to tests: `go test -cpu 1`
* list tests: `go test -list Basics mylib/...`
* run certain tests: `go test -run Basics mylib/...`
* verbose: `go test -v mylib/...`
* run multiple times: `go test -count 10 mylib/...`

## Coverage

The basic usage: `go test -cover mylib`. Per package coverage details are obtained
by using `go test -coverpkg mylib,fmt mylib`.

Generating cover profiles `go test -coverprofile=cover.out mylib`. Then analyse by
using `go tool cover -html=cover.out`.

To figure out how many times a statement was executed during a test run, use `-covermode count`.

## Benchmarking

To test all benchmarking tests in `mylib` run `go test -bench . mylib`.
To test some benchmarking tests in `mylib` run `go test -bench Adding mylib`.
To limit the benchmarking time, use `-benchtime 1s`. See docs for other examples of valid values.
To find out memory allocation issues, add the `-benchmem` param.

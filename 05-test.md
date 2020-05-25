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

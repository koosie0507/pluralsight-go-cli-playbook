# The Go Environment

## Go Version

`go version`

## Environment

`go help env`
`go env [<ENV_VAR1> <ENV_VAR2> ...]`

## Packages

Listing dev env packages should be pretty easy:

* `go help list`
* `go list [-json] [-f {{.SomeProperty}}]`

Take a gander at what properties are available using the help.

Fetching packages is easy, just `go get`. Interesting options:

* install: `go get`
* update: `go get -u`
* download only: `go get -d`
* fix compatibility issues: `go get -fix`

## Looking at Docs

`go help doc` will show you how to use the docs. Very comprehensive, very powerful.

`go doc package/module.method_or_field` is a very nice navigation shortcut. `go doc -u` is
great for getting TOC.

`godoc -http :8000` - shows you golang.org locally.

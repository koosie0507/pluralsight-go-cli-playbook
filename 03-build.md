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

## Build Modes

About: `go help buildmode`

### Shared Libraries

To build shared libraries, first `go build -buildmode=shared std`. Then `go build -buildmode=shared hello`.

To install a shared library `go install -linkshared -buildmode=shared hello`.
Finally, `go build -linkshared cmds/hello`.

You end up with a dynamically linked `hello` executable.

### Exporting to C

Prepare the Go source code, then do either: 

* `go build -buildmode=c-archive src/cmds/hello/hello.go`, _or_
* `go build -buildmode=c-shared src/cmds/hello/hello.go`

### Building Plugins

Have a look in `src/action` and `src/plugin` to see a bare minimum plugin
implementation.

Then `go build -buildmode=plugin -o=plugin.so src/plugin/plugin.go`. Note that
if you don't specify `-o` nothing will be output.

Afterwards, build the app or `go run` it and specify the path to the plugin
via the `-plugin` flag.

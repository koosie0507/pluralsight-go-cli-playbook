# Workspace Commands

## Cleaning workspaces up

`go help clean`

Essentially, run `go clean` with various flags.

## Formatting

`go help fmt`

`go doc gofmt`

Essentially, run `go fmt` with various flags. Useful to get this into git pre-commit hooks.

## Vetting

`go help vet`

`go doc cmd/vet`

Essentially, run `go vet` on the packages you need. Awesome to run this in CI.

## Generating Source Code

`go help generate` - use comments in source code files to generate source code and
make templates.

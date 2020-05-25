# Profiling

## About

* `go help testflag`: `-blockprofile`, `-blockprofilerate`, `-coverprofile`, `-cpuprofile`,
  `-memprofile`, `-memprofilerate`, `-mutexprofile`, `-mutexprofilefraction`

## Visualising

Install Graphviz before doing anything. `go tool` requires it for generating graphs.

`go tool pprof -web <executable> <profiling_info>`

Executable is set to `<your_library_name>.test` whereas `<profiling_info>` is supplied to
one of the `go test` profiling commands as the output name.

### Memory Profiling

Simple execution: `go test -memprofile mem.out mylib && go tool pprof -web mylib.test mem.out`
Smaller sample rates yield more details: `go test -memprofile mem.out -memprofilerate 1 mylib && ...`

### CPU Profiling

Simple execution: `go test -cpuprofile cpu.out mylib && ...`
Control execution to see more details: `go test -cpuprofile cpu.out -count 100000 mylib && ...`

### Tracing

To see what gets executed run the tests in tracing mode.

`go test -trace trace.out mylib && go tool trace trace.out`

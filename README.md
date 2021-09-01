# Profiling in Go

## Profile Types

01. CPU
02. Memory
03. Blocking

## How to enable profiling in go apps

01. Direct use of `runtime/pprof` package
02. Pass flags to go test command
03. Use `net/http/pprof` package


## Download profile/trace files

```bash
curl -o <filepath> <url>
curl -o trace.out http://localhost:6060/debug/pprof/profile

```

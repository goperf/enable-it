# Tracing

## Collect trace

```bash
curl -o trace.out http://localhost:6060/debug/pprof/trace
```

*HTTP GET Parameters*

-seconds - Collect trace for specific duration


## Analyze trace

```bash
go tool trace trace.out
```

It will open a webpage with all the available traces

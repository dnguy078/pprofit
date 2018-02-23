# Pproffit
### Go Profiling and optomization using pprof

---

## Agenda

- Profiling
- How can start profiling our apps 
- Analyzing profiles with benchmarks and pprof features (flame graphs, web, svg)
- Code Demo

---

## pprof
- in built profling tool with many supported profiles:
  - CPU profiles (stack sampling, where your time is being spent)
  - heap (what memory is being allocated and where)
  - all running goroutines and their stack
  - trace, block profiles, garbage collection, etc..


---
## Ways to get started using pprof
- Through tests and benchmarks, using the -cpuprofile and -memprofile flags. Build a profile for a benchmark: 
```
go test . -bench . (or -bench RegexBenchTest) -cpuprofile prof.cpu 
```
Then analyze the profile:
```
pprof [binary] prof.cpu
```

- import _ net/http/pprof to add /debug/pprof endpoints in your service. Start the profile and analyze the results directly:

```
go tool pprof -seconds 5 http://localhost:6060/debug/pprof/profile
```

---?code=main.go&lang=golang&title=main.go

@[7](Expose pprof to the default mux)

---?code=main_test.go&lang=golang&title=main_test.go

@[7](Expose pprof to the default mux)

---?code=Makefile&title=Makefile


---

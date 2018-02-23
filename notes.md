pprof notes

1. pprof
 - start a server exposing pprof endpoint
 - go tool pprof -seconds=10 localhost:8080/debug/pprof/profile  
    ** profile the application for 10 seconds
 - start load testing

2. load test benchmark tool
go get -u github.com/tsliwowicz/go-wrk
** 5 seconds of load testing with 10 goroutines
go-wrk -d 5 http://localhost:8080/dnguyen@golang.org

3. flame graphs go-torch
go-torch -t 5
** runs the same as go tool pprof but produces flame graphs (torch.svg)
- horizontal axis is the time spent while profiling the application
- vertical axis the the stack trace of each function call
- order and color does not irrelevant

4. Optimize notes:
    - move regex must compile out

5. Benchmark
go test -bench .
goos: linux
goarch: amd64
pkg: github.com/dnguy078/pprofit
BenchmarkHandler-4         50000             27855 ns/op
PASS
ok      github.com/dnguy078/pprofit     1.694s

means that the loop ran for 50000 times at a speed of 27855 ns/op
** If a benchmark needs some expensive setup before running, the benchmark time can be reset
ie)
func BenchmarkBigLen(b *testing.B) {
    big := NewBig()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}

- Creates a cpu profile
go test -bench . -cpuprofile prof.cpu
outputs:
    - prof.cpu - cpu profile data
    - pprofit.test benchmark test holds symbolic links to variables?

go get github.com/google/pprof
    pprof -http [host]:[port] [options] [binary] <source> ...
    host and port - to view web interface
    ie) show cpu data on localhost:1234
    pprof -http "localhost:1234" prof.cpu

alternatively you can do
    pprof pprofit.test prof.cpu


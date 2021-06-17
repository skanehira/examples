# go-bench

```sh
MacbookPro13% go test -bench . --benchmem --memprofile mem.pprof
goos: darwin
goarch: arm64
pkg: github.com/skanehira/exmaples/go-bench
BenchmarkF-8        5082            235958 ns/op         4654345 B/op         30 allocs/op
PASS
ok      github.com/skanehira/exmaples/go-bench  2.234s
MacbookPro13% go tool pprof go-bench.test mem.pprof
File: go-bench.test
Type: alloc_space
Time: Jun 17, 2021 at 9:08pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 36.05GB, 100% of 36.05GB total
      flat  flat%   sum%        cum   cum%
   36.05GB   100%   100%    36.05GB   100%  github.com/skanehira/exmaples/go-bench.f (inline)
         0     0%   100%    36.05GB   100%  github.com/skanehira/exmaples/go-bench.BenchmarkF
         0     0%   100%    36.05GB   100%  testing.(*B).launch
         0     0%   100%    36.05GB   100%  testing.(*B).runN
(pprof)
```

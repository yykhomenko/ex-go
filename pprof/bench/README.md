Run:
```bash
GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out
go tool pprof bench.test cpu.out
```

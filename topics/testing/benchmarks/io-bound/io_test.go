package main

import (
	"fmt"
	"runtime"
	"testing"
)

// GOGC=off GOMAXPROCS=1 go test -bench=. -benchtime=3s
// GOGC=off go test -bench=. -benchtime=3s

//Processing 1000 docements using 8 goroutines on 8 threads
//goos: darwin
//goarch: arm64
//pkg: github.com/learning_golang/topics/testing/benchmarks/io-bound
//BenchmarkSequential-8                  3        1366510903 ns/op
//BenchmarkConcurrent-8                 21         166641550 ns/op
//BenchmarkSequentialAgain-8             3        1315832833 ns/op
//BenchmarkConcurrentAgain-8            21         165362804 ns/op

var docs []string

func init() {
	docs = generateDocs(1e3)
	fmt.Printf("Processing %d docements using %d goroutines on %d threads \n", len(docs), runtime.NumCPU(), runtime.GOMAXPROCS(0))
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("Go", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "Go", docs)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("Go", docs)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "Go", docs)
	}
}

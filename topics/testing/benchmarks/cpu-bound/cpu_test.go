package main

// GOGC=off GOMAXPROCS=1 go test -run=none -bench=. -benchtime 3s		// 单个处理线程下，性能基本一样
// GOGC=off go test -run=none -bench . -benchtime 3s	// 多线程处理下，并发性能更高

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var numbers []int

func init() {
	rand.Seed(time.Now().UnixNano())
	numbers = generateNumList(5e8)

	fmt.Printf("Processing %d numbers, using %d goroutines on %d threads \n", len(numbers), runtime.NumCPU(), runtime.GOMAXPROCS(0))
}

// 测试顺序执行
func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

// 测试并发执行
func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

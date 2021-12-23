package example3

// go test -run none -bench . --benchtime 3s --benchmem

//goos: darwin
//goarch: arm64
//pkg: github.com/learning_golang/topics/testing/benchmarks/example3
//BenchmarkSingle-8             66          45744757 ns/op        164379377 B/op   1000001 allocs/op
//BenchmarkNumCPU-8             99          37172131 ns/op        228381372 B/op   5000018 allocs/op
//BenchmarkUnlimited-8           8         417739104 ns/op        235231864 B/op   4062443 allocs/op
// 上面结果的性能排序：NumCPU > Single > Unlimited， 过多的使用goroutine不一定能加速程序运行，需要考虑一下实际负载.

import (
	"testing"
)

var n []int

func init() {
	for i := 0; i < 1_000_000; i++ {
		n = append(n, i)
	}
}

func BenchmarkSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Single(n)
	}
}

func BenchmarkNumCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumCPU(n, 0)
	}
}

func BenchmarkUnlimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unlimited(n)
	}
}

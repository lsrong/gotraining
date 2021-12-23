package example1

// 演示基准测试基础使用.

// go test -run none -bench . --benchtime 3s --benchmem

/**
lsrong@lsrong-Mac basic % go test -run none -bench . --benchtime 3s --benchmem
goos: darwin
goarch: arm64
pkg: github.com/learning_golang/topics/testing/benchmarks/basic
BenchmarkSprint-8       82536538                44.03 ns/op            5 B/op          1 allocs/op
PASS
ok      github.com/learning_golang/topics/testing/benchmarks/basic      6.600s

// 说明:
--benchtime: 基础测试运行时长，默认为1s

BenchmarkName-8: 基础测试运行时对应的GOMAXPROCS的值（8）
82536538: 运行调用被测试代码的测试（b.N = 82536538）
44.03 ns/op: 每次测试代码调用花费的时间， 44.03 微妙
--benchmem: 提供显示每次操作分配内存的次数，以及每次操作分配的字节数的选项。
5 B/op： 每次测试调用的内存分配字节数。
1 allocs/op： 每次测试调用分配内存次数。
*/

import (
	"fmt"
	"testing"
)

var gs string

func BenchmarkSprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gs = s
}

func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello %d", i)
	}

	gs = s
}

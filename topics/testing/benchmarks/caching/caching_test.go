package caching

import "testing"

// 执行命令： go test -run=none -bench=. -benchtime=3s
//Elements in the link list 16777216
//Elements in the matrix 16777216
//goos: darwin
//goarch: arm64
//pkg: github.com/ardanlabs/gotraining/topics/go/testing/benchmarks/caching
//BenchmarkLinkListTraverse-8          206          17444299 ns/op
//BenchmarkColumnTraverse-8             82          44404233 ns/op
//BenchmarkRowTraverse-8               402           8940838 ns/op

var rc int

// BenchmarkLinkedListTraverse 捕获执行 链接列表遍历所需的时间。
func BenchmarkLinkedListTraverse(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = LinkedListTraverse()
	}
	rc = c
}

// BenchmarkColumnTraverse 捕获执行 列遍历所需的时间
func BenchmarkColumnTraverse(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = ColumnTraverse()
	}

	rc = c
}

// BenchmarkRowTraverse 捕获执行 行遍历所需的时间
func BenchmarkRowTraverse(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = RowTraverse()
	}

	rc = c
}

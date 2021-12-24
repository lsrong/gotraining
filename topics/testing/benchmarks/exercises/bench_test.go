package exercises

import (
	"fmt"
	"strconv"
	"testing"
)

// Write three benchmark tests for converting an integer into a string. First using the
// fmt.Sprintf function, then the strconv.FormatInt function and then strconv.Itoa.
// Identify which function performs the best.
// 编写三个将整数转换为字符串的基准测试。
// 首先使用 fmt.Sprintf 函数，然后是 strconv.FormatInt 函数，然后是 strconv.Itoa。确定哪个功能执行得最好。

// go test -run=none -bench=. -benchtime=3s -benchmem
// BenchmarkSprintf-8      84441590                42.92 ns/op            2 B/op          1 allocs/op
// BenchmarkFormat-8       1000000000               2.076 ns/op           0 B/op          0 allocs/op
// BenchmarkItoa-8         1000000000               2.084 ns/op           0 B/op          0 allocs/op

var s string

func BenchmarkSprintf(b *testing.B) {
	var l string
	number := 10
	for i := 0; i < b.N; i++ {
		l = fmt.Sprintf("%d", number)
	}
	s = l
}

func BenchmarkFormatInt(b *testing.B) {
	var l string
	var number int64 = 10
	for i := 0; i < b.N; i++ {
		l = strconv.FormatInt(number, 10)
	}
	s = l
}
func BenchmarkItoa(b *testing.B) {
	var l string
	number := 10
	for i := 0; i < b.N; i++ {
		l = strconv.Itoa(number)
	}
	s = l
}

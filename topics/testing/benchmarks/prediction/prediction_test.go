package prediction

import (
	"math/rand"
	"testing"
	"time"
)

// 可预测和不可预测两种情况的性能区别，
// 演示预测提供的代码来显示分支预测如何影响性能。

// go test -run none -bench . -benchtime 3s -benchmem

var r uint8

func crunch(data []uint8) uint8 {
	var sum uint8
	for _, n := range data {
		if n < 128 {
			sum--
		} else {
			sum++
		}
	}

	return sum
}

// BenchmarkPredictable 运行可测试可以预测的分支代码，data 的值都是0
func BenchmarkPredictable(b *testing.B) {
	data := make([]uint8, 1024)
	b.ResetTimer()

	var a uint8

	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	r = a
}

// BenchmarkUnpredictable 运行不可预测的分支代码， data 的值是随机值
func BenchmarkUnpredictable(b *testing.B) {
	data := make([]uint8, 1024)
	rand.Seed(time.Now().UnixNano())
	for i := range data {
		data[i] = uint8(rand.Uint32())
	}

	b.ResetTimer()
	var a uint8
	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	r = a
}

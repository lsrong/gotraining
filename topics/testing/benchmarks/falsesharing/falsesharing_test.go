package falsesharing

// false sharing
// 伪共享的非标准定义为：缓存系统中是以缓存行（cache line）为单位存储的，当多线程修改互相独立的变量时，如果这些变量共享同一个缓存行，就会无意中影响彼此的性能，这就是伪共享。

import (
	"sync"
	"testing"
)

// 测试显示错误共享对并发内存写入的影响。
//go test -run=none -bench=. -benchtime=3s --benchmem
//goos: darwin
//goarch: arm64
//pkg: github.com/learning_golang/topics/testing/benchmarks/falsesharing
//BenchmarkGlobal-8            805           4314753 ns/op              65 B/op          0 allocs/op
//BenchmarkLocal-8            2257           1614284 ns/op              11 B/op          0 allocs/op

type cnt struct {
	counter int64
}

const grs = 8

var countersPad [grs]cnt

// BenchmarkGlobal 测试了 8 个 goroutine 并行递增全局计数器的性能。（访问全局会出现共享情况）
func BenchmarkGlobal(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(grs)
		for g := 0; g < grs; g++ {
			go func(i int) {
				for {
					countersPad[i].counter++

					// 结束计数
					if countersPad[i].counter%1e6 == 0 {
						wg.Done()
						return
					}
				}
			}(g)
		}
		wg.Wait()
	}
}

// BenchmarkLocal 测试了 8 个 goroutine 增加其本地计数器的性能。（本地计数则不会共享全局计数）
func BenchmarkLocal(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(grs)
		for g := 0; g < grs; g++ {
			go func(i int) {
				var counter int64
				for {
					// 利用局部变量计数
					counter++

					if counter%1e6 == 0 {
						// 将最终计数器写入特定的全局计数器。
						countersPad[i].counter = counter
						wg.Done()
						return
					}
				}
			}(g)
		}
		wg.Wait()
	}
}

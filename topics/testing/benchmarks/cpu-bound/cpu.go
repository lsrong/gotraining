package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//  了解您的工作负载对于了解是否可以并发执行以及执行起来有多复杂至关重要.

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	numbers := generateNumList(1e6)
	fmt.Printf("add numbers: %d \n", add(numbers))
	fmt.Printf("addConcurrent numbers: %d\n", addConcurrent(runtime.NumCPU(), numbers))
}

// 创建随机的数值元素的切片
func generateNumList(totalNumber int) []int {
	numbers := make([]int, totalNumber)
	for i := 0; i < totalNumber; i++ {
		numbers[i] = rand.Intn(totalNumber)
	}

	return numbers
}

// 顺序统计所有元素的总数
func add(numbers []int) int {
	var counter int
	for _, n := range numbers {
		counter += n
	}

	return counter
}

// 并发统计所有元素的总数
func addConcurrent(goroutines int, numbers []int) int {
	var counter int64
	totalNumber := len(numbers)
	step := totalNumber / goroutines
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		go func(g int) {
			start := (g - 1) * step
			end := start + step
			if g == goroutines {
				end = totalNumber
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}
			atomic.AddInt64(&counter, int64(lv))
			wg.Done()
		}(i)
	}

	wg.Wait()

	return int(counter)
}

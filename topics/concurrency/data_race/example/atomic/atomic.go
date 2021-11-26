package main

// Sample program to show how to use the atomic package to
// provide safe access to numeric types.

// atomic 提供原子操作，安全操作连接整型

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var wg sync.WaitGroup

func main() {
	const grs = 2
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		routine()
	}

	wg.Wait()
	fmt.Printf("Should be 4, Result is %d", counter)
}

func routine() {
	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Millisecond)
			// atomic
			atomic.AddInt64(&counter, 1)
		}

		wg.Done()
	}()
}

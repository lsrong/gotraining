package main

// Sample program to show how to use a mutex to define critical
// sections of code that need synchronous access.

// 使用互斥锁定义需要同步访问的代码的关键部分.

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

// sync.Mutex
var mutex sync.Mutex

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
			//  只允许一个goroutine 通过互斥锁临界区。
			mutex.Lock()
			{
				value := counter
				time.Sleep(time.Nanosecond)
				value++
				counter = value
			}
			mutex.Unlock()
		}
		wg.Done()
	}()
}

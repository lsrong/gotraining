package main

// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
// 此示例程序演示了如何使用工作包来使用 goroutine 池来完成工作。

import (
	"fmt"
	"sync"
	"time"

	"github.com/learning_golang/topics/concurrency/patterns/task"
)

var names = []string{
	"Li",
	"Ming",
	"Ye",
	"Liu",
}

// namePrinter implements the task.Worker interfaces.
type namePrinter struct {
	name string
}

func (np namePrinter) Work() {
	fmt.Println(np.name)
	time.Sleep(3 * time.Second)
}

func main() {
	const grs = 2

	// Crate a task pool
	tm := task.New(grs)

	var wg sync.WaitGroup
	wg.Add(grs * len(names))
	// 演示循环模拟处理发起处理任务
	for i := 0; i < grs; i++ {
		for _, n := range names {
			np := namePrinter{
				name: n,
			}

			go func() {
				// 提交要处理的任务。当 Do 返回时，它正在被处理。
				tm.Do(np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	tm.Shutdown()
	fmt.Println("shutdown task.")
}

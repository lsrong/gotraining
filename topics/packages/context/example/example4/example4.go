package main

import (
	"context"
	"fmt"
	"time"
)

// Sample program to show how to use the WithTimeout function of the Context package.
// 演示如何使用 Context 包的 WithTimeout 函数的示例程序。

/**
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout))。

*/

type data struct {
	UserID string
}

func main() {
	demoWithTimeout()

	demoWithTimeoutWork()
}

// demoWithTimeoutWork 模拟执行长时间超时任务
func demoWithTimeoutWork() {
	duration := 50 * time.Millisecond

	// 创建具有超时时间的上下文
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// 模拟异步工作任务
	ch := make(chan data, 1)
	go func() {
		// 模拟执行长时间任务
		//time.Sleep(10 * time.Millisecond)
		time.Sleep(100 * time.Millisecond)

		ch <- data{"1"}
	}()

	// 等待任务执行完成
	select {
	case d := <-ch:
		fmt.Println("work complete: ", d)
	case <-ctx.Done():
		fmt.Println("work canceled")
	}

}

// demoWithTimeout 官方demo
func demoWithTimeout() {
	duration := 50 * time.Millisecond

	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		// print: context deadline exceeded
		fmt.Println(ctx.Err())
	}
}

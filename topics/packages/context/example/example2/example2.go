package main

import (
	"context"
	"fmt"
	"time"
)

// Sample program to show how to use the WithCancel function.
// 演示如何使用 WithCancel 函数的示例程序。

/**
context.WithCancel
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

WithCancel returns a copy of parent with a new Done channel. The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.
WithCancel 返回具有新 Done 通道的 parent 副本。返回的上下文的完成通道在调用返回的取消函数或父上下文的完成通道关闭时关闭，以先发生者为准。

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.
取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用取消
*/

func main() {
	// 创建一个只能手动取消的上下文。无论结果如何，都必须调用取消函数。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		// 等待任务完成,如果超过时间就放弃继续执行
		case <-time.After(100 * time.Millisecond):
			fmt.Println("moving on")
		case <-ctx.Done():
			fmt.Println("work complete")
		}
	}()

	// 模拟任务.
	time.Sleep(50 * time.Millisecond)

	// 通知任务已完成
	cancel()

	time.Sleep(time.Second)
}

package main

import (
	"context"
	"fmt"
	"time"
)

// Sample program to show how to use the WithDeadline function.
// 演示如何使用 WithDeadline 函数的示例程序。

/**
context.WithDeadline:

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d.
If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent.
The returned context's Done channel is closed when the deadline expires, when the returned cancel function is called,
or when the parent context's Done channel is closed, whichever happens first.
WithDeadline 返回父上下文的副本，截止日期调整为不迟于 d。如果父节点的截止日期已经早于 d，
WithDeadline(parent, d) 在语义上等价于父节点。返回的上下文的 Done 通道在截止日期到期、
调用返回的取消函数或父上下文的 Done 通道关闭时关闭，以先发生者为准。

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.
取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用取消。
*/

type data struct {
	UserID string
}

func main() {
	// 设置一个超时时期
	deadline := time.Now().Add(150 * time.Millisecond)

	// context.WithDeadline: 生成可以手动取消以及设定一个超时时限的上下文
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// 模拟业务逻辑执行
	ch := make(chan data, 1)

	// 启动一个goroutine模拟进行异步操作
	go func() {
		// 模拟进行执行, 时间超过之前制定的超时时间
		//time.Sleep(100 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)

		// 上报结果
		ch <- data{"1"}
	}()

	// 等待异步工作完成,如果花费过长时间则停止等待
	select {
	case d := <-ch:
		fmt.Println("work complete: ", d)
	case <-ctx.Done():
		fmt.Println("work canceled")
	}
}

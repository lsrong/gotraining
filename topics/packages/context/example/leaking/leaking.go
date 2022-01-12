package main

import (
	"context"
	"fmt"
	"time"
)

// leakGen 一旦调用者完成了生成器（当它中断循环时），goroutine 将永远运行并执行无限循环。我们的代码会泄漏一个 goroutine。
func leakGen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()

	return ch
}

// gen 生成器可以在上下文的 Done 通道上进行选择，一旦上下文完成，就可以取消内部 goroutine
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Gen is closed")
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func main() {
	// 上下文方式避免goroutine泄露
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			cancel()
			break
		}
	}

	time.Sleep(time.Second)
}

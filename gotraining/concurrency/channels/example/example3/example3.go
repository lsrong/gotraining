package main

import (
	"fmt"
	"sync"
	"time"
)

// Sample program to show how to use an unbuffered channel to
// simulate a relay race between four goroutines.

// 演示如何使用无缓冲通道模拟四个 goroutine 之间的接力竞赛的示例程序。

var wg sync.WaitGroup

func main() {
	stack := make(chan int)
	wg.Add(1)
	go Runner(stack)
	stack <- 1
	wg.Wait()
}

func Runner(stack chan int) {
	const maxExchanges = 4
	var exchange int

	// 核心： 无缓冲通道，会阻塞等待到这里，指到第一个goroutine运行到stack <- exchange
	baton := <-stack
	fmt.Printf("Runner %d Running with Baton \n", baton)

	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("Runner %d to the Line \n", exchange)
		go Runner(stack)
	}

	time.Sleep(100 * time.Millisecond)

	// 交互次数为4时则结束流程.
	if baton == maxExchanges {
		fmt.Printf("Runner %d Finish,Race Over \n", baton)
		wg.Done()
		return
	}

	fmt.Printf("Runer %d Exchange With Runner %d \n", baton, exchange)

	stack <- exchange
}

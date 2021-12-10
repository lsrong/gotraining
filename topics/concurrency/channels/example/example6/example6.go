package main

// 演示使用goroutine 顺序打印数字: 0,1,2,3,4,5,6,7,8,9

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	//go display1(n)

	var wg sync.WaitGroup
	counter := make(chan int)
	wg.Add(2)
	go func() {
		display2(n, counter)
		wg.Done()
	}()

	go func() {
		display2(n, counter)
		wg.Done()
	}()
	counter <- 0
	wg.Wait()
}

// display1 不能实现
func display1(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

// display2 使用无缓冲通道开启两个 goroutine 之间来回交换打印叠加数字的思路
func display2(n int, counter chan int) {
	for {
		c, ok := <-counter
		if !ok {
			return
		}

		if c == n {
			close(counter)
			return
		}
		fmt.Printf("counter: %d \n", c)
		c++
		counter <- c
	}
}

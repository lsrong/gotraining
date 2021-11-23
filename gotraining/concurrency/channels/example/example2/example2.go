package main

// Sample program to show how to use an unbuffered channel to
// simulate a game of tennis between two goroutines.
// 使用无缓冲通道模拟两个 goroutine 之间的网球比赛的示例程序。

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动两个玩家
	go func() {
		player("Li", court)
		wg.Done()
	}()

	go func() {
		player("Ye", court)
		wg.Done()
	}()

	// start to set
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	for {
		ball, wd := <-court
		if !wd {
			// 一旦channel关闭说明上一个loop失败，则当前为赢。
			fmt.Printf("Player %s Won \n", name)
			return
		}
		n := rand.Intn(100)
		if n%7 == 0 {
			fmt.Printf("Player %s Missed \n", name)

			// Close 关闭channel,发送失败的信号
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		court <- ball

	}
}

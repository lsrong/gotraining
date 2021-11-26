package main

// This sample program demonstrates how the logger package works.

// 这个示例程序演示了 logger 包是如何工作的。

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/learning_golang/gotraining/concurrency/patterns/logger"
)

// device 模拟一个设备写入日志信息
type device struct {
	problem bool
}

// Write 实现io.Writer接口
func (d *device) Write(data []byte) (n int, err error) {
	for d.problem {

		// 模拟磁盘问题。即不能写入日志
		time.Sleep(time.Second)
	}
	fmt.Print(string(data))

	return len(data), nil
}

func main() {
	var d device
	const grs = 10
	l := logger.New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// 要控制模拟磁盘阻塞。捕获中断信号以切换设备问题.
	for {
		<-sigCh

		d.problem = !d.problem
	}
}

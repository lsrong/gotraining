package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.

// 此示例程序演示了如何使用通道来监视程序运行的时间量并在程序运行时间过长时终止该程序。

const timeoutSecond = 3 * time.Second

func main() {
	// 系统中断信号
	sigChan := make(chan os.Signal, 1)

	// 超时信号
	timeout := time.After(timeoutSecond)

	// 是否完成信号
	complete := make(chan error)

	// 关闭信号
	shutdown := make(chan struct{})

	// 监听系统中断信号
	signal.Notify(sigChan, os.Interrupt)

	log.Println("Starting Process")
	go processor(complete, shutdown)
ControlLoop:
	for {
		select {
		case <-sigChan:
			// 系统中断信号
			log.Println("OS INTERRUPT")

			// 关闭信号,通知中程序需要关闭服务
			close(shutdown)
			// 停止接受其他信号
			sigChan = nil
		case <-timeout:
			// 超时
			log.Println("Timeout - Killing Program")
			os.Exit(1)
		case err := <-complete:
			log.Printf("Task Completed: Error [%s]", err.Error())
			//break // break只能停止select , 不能停止for
			break ControlLoop
		}
	}

}

func processor(complete chan<- error, shutdown <-chan struct{}) {
	log.Println("Starting ...")

	var err error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic: ", r)
		}
		complete <- err
	}()

	err = doWork(shutdown)

	log.Println("Completed!")
}

func doWork(shutdown <-chan struct{}) error {
	log.Println("Processor - Task 1")
	time.Sleep(time.Second * 2)
	if checkShutdown(shutdown) {
		return errors.New("Early shutdown ")
	}

	log.Println("Processor - Task 2")
	time.Sleep(time.Second * 1)
	if checkShutdown(shutdown) {
		return errors.New("Early shutdown ")
	}
	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}

func checkShutdown(shutdown <-chan struct{}) bool {
	select {
	case <-shutdown:
		// 执行关闭操作
		log.Println("check shutdown Early ")
		return true
	default:
		return false
	}
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
// 这个示例程序演示了 goroutine 信号的基本通道机制。
// 专注于简化所需编排的信令和语义

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 父协程接受子协程发送的信号任务
	//waitForResult()
	//fanOut()
	//fanOutSem()

	// 父协程发送信号，子协程负责处理任务
	//waitForTask()
	//pooling()
	//boundedWorkPooling()

	// 丢弃来不及处理的任务，select： 选择其中一个case
	//drop()

	// context.Timeout  + select 实现超时任务模式
	//cancellation()

	// 实现简单的重试任务
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//retryTimeout(ctx, time.Second, func(ctx context.Context) error { return errors.New("Failed ") })

	// 利用channel+context实现可取消操作任务 ?
	stop := make(chan struct{})
	channelCancellation(stop)
}

// waitForResult 父goroutine等待子goroutine发送完成任务结果信息号
func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child: send signal")
	}()
	time.Sleep(time.Second)
	result := <-ch

	fmt.Println("parent: recv'd signal : ", result)
	fmt.Println("----------------")
}

// fanOut 父goroutine创建足够多的子goroutine处理任务，并等待任务完成.
func fanOut() {
	children := 10
	ch := make(chan string, children)
	for i := 0; i < children; i++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child: send signal: ", child)
		}(i)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println("parent: recv'd signal: ", d)
	}

	time.Sleep(time.Second)
	fmt.Println("----------------")
}

// fanOutSem：将信号量添加到有缓存模式以限制可以调度运行的子 goroutine 的数量。
func fanOutSem() {
	childs := 20
	ch := make(chan string, childs)
	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)
	for c := 0; c < childs; c++ {
		go func(child int) {
			sem <- true
			{
				ch <- "data"
				fmt.Printf("child send signal: %d \n", child)
			}
			<-sem
		}(c)
	}

	for childs > 0 {
		data := <-ch
		childs--
		fmt.Println("parent recv'd signal: ", data)
	}

	time.Sleep(time.Second)
	fmt.Println("----------------")
}

//waitForTask  父goroutine发送信号，子goroutine等待接受通道信号.
func waitForTask() {
	ch := make(chan string)
	go func() {
		data := <-ch
		fmt.Println("child: recv'd signal: ", data)
	}()
	ch <- "data"
	fmt.Println("parent send signal")
	fmt.Println("----------------")
}

// pooling 在这种模式中，父协程向等待工作执行的子协程池发出 10 个工作信号。
func pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for i := 0; i < g; i++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d: recv'd: %s \n", child, d)
			}
			fmt.Printf("child %d shutdown signal \n", child)
		}(i)
	}
	const worker = 10
	for i := 0; i < worker; i++ {
		ch <- "data"
		fmt.Println("parent: send signal: ", i)
	}

	// close 已关闭的channel不允许发送，但是允许接口channel内的数据，信号
	close(ch)
	time.Sleep(time.Second)

	fmt.Println("----------------")
}

// boundedWorkPooling：在这种模式中，
// 创建了一个子 goroutine 池来为固定数量的工作提供服务。
// 父 goroutine 迭代所有工作，将其发送到池中。
// 一旦所有工作都发出信号，然后关闭通道，刷新通道，并且子 goroutine 终止。
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)
	ch := make(chan string, g)
	for i := 0; i < g; i++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("Child %d : recv'd signal: %s \n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(i)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
	fmt.Println("----------------")
}

// drop：在这种模式中，父 goroutine 向无法处理所有工作的单个子 goroutine 发出 2000 件工作信号。
// 如果父级执行发送而子级未准备好，则该工作将被丢弃并丢弃
func drop() {
	const cap = 10
	ch := make(chan string, cap)
	go func() {
		for p := range ch {
			fmt.Println("child: recv'd signal: ", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent: send signal: ", w)
		default:
			fmt.Println("parent: dropped work: ", w)
		}
	}
	close(ch)

	time.Sleep(time.Second)
	fmt.Println("----------------")
}

// cancellation：在这种模式中，父 goroutine 创建一个子 goroutine 来执行一些工作。
// 父 goroutine 只愿意等待 150 毫秒才能完成该工作。
// 150 毫秒后，父 goroutine 走开。
func cancellation() {
	duration := 150 * time.Millisecond
	// context的WithTimeout 实现超时任务取消的小示例
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
	}()

	// select: 同一时间只会选中其中一个case
	select {
	case d := <-ch:
		fmt.Println("work completed", d)
	case <-ctx.Done():
		fmt.Println("work canceled")
	}
	fmt.Println("----------------")
}

// retryTimeout：需要验证是否可以在没有错误的情况下完成某些操作，但在此之前可能需要一些时间。
//	您设置重试间隔以在重试调用之前创建延迟，并使用上下文设置超时。
func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {
	for {
		if err := check(ctx); err == nil {
			fmt.Println("Work finished successfully")
			return
		}

		// check if timeout has expired
		if ctx.Err() != nil {
			fmt.Println("time expired 1 : ", ctx.Err())
			return
		}

		// retry
		t := time.NewTimer(retryInterval)
		select {
		case <-ctx.Done():
			fmt.Println("time expired 2: ", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

// channelCancellation shows how you can take an existing channel being
// used for cancellation and convert that into using a context where
// a context is needed.

// channelCancellation 如何获取用于取消的现有通道并将其转换为使用需要上下文的上下文。
func channelCancellation(stop <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case _ = <-stop: // if a signal is received on the stop channel , cancel the context.
			cancel()
			fmt.Println("cancel")
		case <-ctx.Done():
		}
	}()

	// Imagine a function that is performing an I/O operaion that is cancelable
	err := func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.baidu.com", nil)
		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		return nil
	}(ctx)

	fmt.Println(err)
}

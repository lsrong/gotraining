// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.

// 编写一个程序，创建一组固定的工人来生成随机数。
// 丢弃任何可被 2 整除的数字。继续接收直到接收到 100 个数字。在终止之前告诉工人关闭。

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Add imports.

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	var workers = runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(workers)

	// Create a fixed size pool of goroutines to generate random numbers.
	for w := 0; w < workers; w++ {
		go func() {
			for {
				number := rand.Intn(1000)
				select {
				case values <- number:
				case <-shutdown:
					wg.Done()
					return
				}
			}
		}()
	}

	// Create a slice to hold the random numbers.
	var ints []int

	// Receive from the values channel with range.
	for n := range values {

		// continue the loop if the value was even.
		if n%2 == 0 {
			continue
		}
		// Store the odd number.
		ints = append(ints, n)

		// break the loop once we have 100 results.
		if len(ints) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	fmt.Println(len(ints), ints)
}

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.

// 编写一个程序，其中两个 goroutine 来回传递一个整数十次。
// 当每个 goroutine 接收到整数时显示。每次通过增加整数。
// 一旦整数等于 10，干净地终止程序。
package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	court := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(1)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine(court)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine(court)
		wg.Done()
	}()

	// Send a value to start the counting.
	court <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(court chan int) {
	for {
		// Wait for the value to be sent.
		n, ok := <-court
		// If the channel was closed, return.
		if !ok {
			return
		}

		// Display the value.
		fmt.Println(n)

		// Terminate when the value is 10.
		if n == 10 {
			close(court)
			return
		}

		// Increment the value and send it
		n++
		// over the channel.必须有下一个goroutine准备接受通道court
		court <- n
	}
}

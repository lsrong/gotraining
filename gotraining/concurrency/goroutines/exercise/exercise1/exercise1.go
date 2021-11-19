package main

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	var count = 100

	wg.Add(2)
	fmt.Println("Create goroutines")
	go func() {
		for i := count; i >= 0; i-- {
			fmt.Printf("[A:%d]\n", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= count; i++ {
			fmt.Printf("[B:%d]\n", i)
		}
		wg.Done()
	}()

	fmt.Println("Waiting to finish")

	wg.Wait()

	fmt.Println("\n Terminating")
}

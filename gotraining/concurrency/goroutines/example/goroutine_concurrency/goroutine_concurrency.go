package main

// Sample program to show how to create goroutines and
// how the scheduler behaves.

// 演示创建goroutines和调度行为

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
	wg.Add(2)
	fmt.Println("Start goroutines")

	// uppercase loop
	go func() {
		uppercase()
		wg.Done()
	}()

	// lowercase loop
	go func() {
		lowercase()
		wg.Done()
	}()

	fmt.Println("Waiting to Finish")

	wg.Wait()
	fmt.Println("\n Terminating program")

}

func uppercase() {
	for count := 0; count < 3; count++ {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c ", i)
		}
	}
}

func lowercase() {
	for count := 0; count < 3; count++ {
		for i := 'a'; i <= 'z'; i++ {
			fmt.Printf("%c ", i)
		}
	}
}

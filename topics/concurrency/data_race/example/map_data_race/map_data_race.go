package main

import (
	"fmt"
	"sync"
)

// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.

// 演示默认的map类型是并发不安全的。

var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			scores["A"]++
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			scores["B"]++
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Scores:", scores)

	// fatal error: concurrent map writes
}

package main

// $ go build -race

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.

// 演示出现竞态的场景，需要避免此类情况。

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func main() {
	const grs = 2
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		routine()
	}
	wg.Wait()
	fmt.Printf("Should be 4, Result is %d", counter)
}

func routine() {
	go func() {
		for i := 0; i < 2; i++ {
			value := counter
			time.Sleep(time.Nanosecond)
			value++
			counter = value
		}
		wg.Done()
	}()
}

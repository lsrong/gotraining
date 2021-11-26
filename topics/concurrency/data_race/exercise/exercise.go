package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var numbers []int

var mutex sync.Mutex

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const grs = 3
	var wg sync.WaitGroup

	wg.Add(grs)

	// Create three goroutines to generate random numbers
	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	wg.Wait()

	for i, n := range numbers {
		fmt.Println(i, n)
	}
}

// random generates random numbers and stores them into a slice
func random(counter int) {
	for i := 0; i < counter; i++ {
		n := rand.Intn(100)
		// Multiple Goroutines will cause race conditions, and atomic operation must be used.
		mutex.Lock()
		{
			numbers = append(numbers, n)
		}
		mutex.Unlock()
	}
}

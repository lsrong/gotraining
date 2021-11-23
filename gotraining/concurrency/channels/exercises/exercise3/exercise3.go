// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Declare constant for number of goroutines.
const goroutines = 100

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	nums := make(chan int, goroutines)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Iterate and launch each goroutine.
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			n := rand.Intn(1000)
			if n%2 == 0 {
				return
			}
			nums <- n

		}()
	}

	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.
	go func() {
		wg.Wait()
		close(nums)
	}()

	// Receive from the channel until it is closed.
	// Store values in a slice of ints.
	var ints []int
	for n := range nums {
		ints = append(ints, n)
	}
	// Print the values in our slice.
	fmt.Println(ints)
}

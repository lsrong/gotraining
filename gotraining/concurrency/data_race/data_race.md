## Data Races

A data race is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that appear totally random or can never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

数据竞争是指两个或多个 goroutine 尝试同时读取和写入同一资源。竞争条件可能会产生完全随机的错误，或者在破坏数据时永远不会出现。原子函数和互斥锁是一种同步 goroutine 之间共享资源访问的方法。

## Notes

* Goroutines need to be coordinated and synchronized.
* Goroutines 需要协调和同步.
* 
* When two or more goroutines attempt to access the same resource, we have a data race.
* 当两个或多个 goroutine 试图访问同一个资源时，我们就会发生数据竞争。
*
* Atomic functions and mutexes can provide the support we need.
* 原子函数和互斥锁可以提供我们需要的支持。

## Cache Coherency and False Sharing
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (30:09-38:30)](https://youtu.be/WDIkqP4JbkE?t=1809)  
[Code Example](../../testing/benchmarks/falseshare/README.md)

![figure1](figure1.png)

## Cache Coherency and False Sharing Notes 缓存一致性和错误共享说明

* Thread memory access matters. 
* 线程内存访问很重要。
* If your algorithm is not scaling look for false sharing problems.
* 如果您的算法未扩展，请查找错误共享问题。

## Links

[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206) - Herb Sutter    
[The Go Memory Model](https://golang.org/ref/mem)    
[Introducing the Go Race Detector](http://blog.golang.org/race-detector) - Dmitry Vyukov and Andrew Gerrand    
[Detecting Race Conditions With Go](https://www.ardanlabs.com/blog/2013/09/detecting-race-conditions-with-go.html) - William Kennedy    
[Data Race Detector](https://golang.org/doc/articles/race_detector.html)

## Diagram

### View of Data Race in Example1.

![Data_Race](data_race.png)

## Code Review

[Data Race](example/data_race/data_race.go) ([Go Playground](https://play.golang.org/p/czqXM5wOspX))    
[Atomic Increments](example/atomic/atomic.go) ([Go Playground](https://play.golang.org/p/5ZtLaX7zxt7))    
[Mutex](example/mutex/mutex.go) ([Go Playground](https://play.golang.org/p/ZKE2v9H4oS-))    
[Read/Write Mutex](example/rwmutex/rwmutex.go) ([Go Playground](https://play.golang.org/p/-iXzElPBnDM))    
[Map Data Race](example/map_data_race/map_data_race.go) ([Go Playground](https://play.golang.org/p/ktWRjcJWNjw))

## Advanced Code Review

[Interface Based Race Condition](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/fwRTeBQrZVW))

## Exercises

### Exercise 1
Given the following program, use the race detector to find and correct the data race.

	// https://play.golang.org/p/F5DCJTZ6Lm

	// Fix the race condition in this program.
	package main

	import (
		"fmt"
		"math/rand"
		"sync"
		"time"
	)

	// numbers maintains a set of random numbers.
	var numbers []int

	// init is called prior to main.
	func init() {
		rand.Seed(time.Now().UnixNano())
	}

	// main is the entry point for the application.
	func main() {
		// Number of goroutines to use.
		const grs = 3

		// wg is used to manage concurrency.
		var wg sync.WaitGroup
		wg.Add(grs)

		// Create three goroutines to generate random numbers.
		for i := 0; i < grs; i++ {
			go func() {
				random(10)
				wg.Done()
			}()
		}

		// Wait for all the goroutines to finish.
		wg.Wait()

		// Display the set of random numbers.
		for i, number := range numbers {
			fmt.Println(i, number)
		}
	}

	// random generates random numbers and stores them into a slice.
	func random(amount int) {
		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}
	}

[Template](template/template.go) ([Go Playground](https://play.golang.org/p/Mzt11_xe_ou)) |
[Answer](exercise/exercise.go) ([Go Playground](https://play.golang.org/p/KAakUVF_1k-))

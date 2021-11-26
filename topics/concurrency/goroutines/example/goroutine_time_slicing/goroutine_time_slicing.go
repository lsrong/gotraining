package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// $ go build -o example.exe
// $ ./example.exe | cut -c1 | grep '[AB]' | uniq

// 演示goroutines调度器如何在单线程上切分时间。

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	gr := 2
	wg.Add(gr)
	fmt.Println("Create goroutine")

	// The First goroutine
	go func() {
		printHash("A")
		wg.Done()
	}()

	// The Second goroutine
	go func() {
		printHash("B")
		wg.Done()
	}()

	fmt.Println("Waiting to finish!")
	wg.Wait()

	fmt.Println("\n Terminating program")
}

// printHash 简单打印数字的哈希值
func printHash(prefix string) {
	for i := 0; i < 50000; i++ {
		num := strconv.Itoa(i)

		shaSum := sha1.Sum([]byte(num))

		fmt.Printf("%s: %05d : %x \n", prefix, i, shaSum)
	}
}

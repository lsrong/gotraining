package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Loop: ", i)
		}
	}()

}

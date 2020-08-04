package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world  ", i, " times")
	}
}

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

	// 变量命令规范
	
	for i:=2;i<= 4 ;i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello ", i , "times")
	}
	
}

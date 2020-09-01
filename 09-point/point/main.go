package main

import "fmt"

func initPoint() {
	// 地址：&
	var a int
	a = 100
	fmt.Printf("a address is %p, a=%d", &a, a)

	// 引用定义
}

func main() {
	initPoint()
}

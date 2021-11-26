package main

import "fmt"

// Sample program to show how anonymous functions and closures work.

// 匿名函数与闭包

func main() {
	var n int

	// 定义匿名函数并调用
	func() {
		fmt.Println("Direct:", n)
	}()

	// 匿名函数赋值给变量
	f := func() {
		fmt.Println("variable:", n)
	}

	// 变量调用匿名函数
	f()

	// defer 调用匿名函数
	defer func() {
		fmt.Println("Defer 1 :", n)
	}()

	n = 3
	f()

	defer func() {
		fmt.Println("Defer 2 :", n)
	}()
}

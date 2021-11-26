package main

import (
	"fmt"
	"time"
)

// Sample program to show how literal, constant and variables work
// within the scope of implicit conversion.

// 文字、常量和变量如何在隐式转换范围内工作

func main() {
	now := time.Now()

	// 使用文字常量从现在减去 5 纳秒。
	literal := now.Add(-5)

	// 使用声明的常量从现在减去 5 秒。
	const timeout = 5 * time.Second
	constant := now.Add(-timeout)

	// 使用 int64 类型的变量从现在减去 5 纳秒。
	minusFive := int64(5)
	// ./implicit-conversion.go:24:22: cannot use -minusFive (type int64) as type time.Duration in argument to now.Add
	// variable := now.Add(-minusFive)
	// 变量不支持隐式表达式类型转换,需要显式转换类型
	variable := now.Add(-time.Duration(minusFive))

	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Literal : %v\n", literal)
	fmt.Printf("Constant: %v\n", constant)
	fmt.Printf("Variable: %v\n", variable)
}

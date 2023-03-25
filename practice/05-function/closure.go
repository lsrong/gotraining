package main

import (
	"fmt"
	"strings"
	"time"
)

/**
闭包的概念：是可以包含自由（未绑定到特定对象）变量的代码块，
这些变量不在这个代码块内或者任何全局上下文中定义，
而是在定义代码块的环境中定义。
要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）
为自由变量提供绑定的计算环境（作用域）。

闭包的价值 : 闭包的价值在于可以作为函数对象或者匿名函数，
对于类型系统而言，这意味着不仅要表示数据还要表示代码。
支持闭包的多数语言都将函数作为第一级对象，
就是说这些函数可以存储到变量中作为参数传递给其他函数，
最重要的是能够被函数动态创建和返回。

Go语言中的闭包同样也会引用到函数外的变量。
闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。
*/
/*
func main() {
	c1 := Addr()
	fmt.Printf("闭包Addr示例1:%d, %d \n", c1(10), c1(100))
	c2 := Addr()
	fmt.Printf("闭包Addr示例2:%d, %d \n", c2(20), c2(200))

	c3 := add(100)
	fmt.Printf("闭包add示例1:%d, %d \n", c3(10), c3(100))
	c4 := add(200)
	fmt.Printf("闭包add示例2:%d, %d \n", c4(10), c4(100))

	c5 := makeSuffix(".png")
	fmt.Printf("闭包makeSuffix示例: %s \n", c5("demo"))

	c6, c7 := manyClosureFunc(100)
	fmt.Printf("闭包manyClosureFunc示例: %d, %d \n", c6(10), c7(10))

	closureAnonymous()
}
*/

// Addr 简单的闭包，返回值为函数变量
func Addr() func(int) int {
	var x int // 局部变量,零值为0
	return func(y int) int {
		return x + y
	}
}

// add 相加的闭包
func add(a int) func(int) int {
	return func(i int) int {
		a += i
		return a
	}
}

// makeSuffix 一个增加文件后缀名的闭包函数
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

// manyClosureFunc 定义多个闭包的处理函数
func manyClosureFunc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// closureAnonymous 使用匿名函数
func closureAnonymous() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Printf("匿名函数【closureAnonymous】示例: %d \n", index)
		}(i)
		time.Sleep(time.Second)
	}
}

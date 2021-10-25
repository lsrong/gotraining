package main

import (
	"fmt"
	"strings"
	"time"
)

/**
闭包的概念：是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。

闭包的价值 : 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。

Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。
*/

/* 定义一个简单的闭包 */
func Addr() func(int) int {
	var x int // 用户函数局部变量
	return func(d int) int {
		x += d
		return x
	}
}
func testClosureAddr() {
	// 闭包使用
	closure1 := Addr() // 注意闭包函数的变量所用域
	fmt.Println(closure1(10), closure1(100))
	closure2 := Addr()
	fmt.Println(closure2(20), closure2(200))
}

/* 使用闭包使用函数参数变量 */
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
func testClosureAdd() {
	closure1 := add(100)
	fmt.Println(closure1(10), closure1(100))

	closure2 := add(200)
	fmt.Println(closure2(10), closure2(100))
}

/* 定义一个增加文件后缀名的闭包函数 */
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}
func testMakeSuffix() {
	closure := makeSuffix(".png")
	fmt.Println(closure("test"))
}

/* 定义多个闭包处理函数 */
func manyClosureFunc(base int) (func(int) int, func(int) int) {
	// 使用相同的base变量
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
func testManyClosureFunc() {
	f1, f2 := manyClosureFunc(10)
	fmt.Println(f1(1), f2(1))
	fmt.Println(f1(2), f2(2))
	fmt.Println(f1(3), f2(3))
}

/* 使用匿名函数处理：go func(){} () */
func closureAnonymous() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
		time.Sleep(time.Second)
	}

}

func main() {
	testClosureAddr()

	testClosureAdd()

	testMakeSuffix()

	testManyClosureFunc()

	closureAnonymous()

}

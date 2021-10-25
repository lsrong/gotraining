package main

import "fmt"

/**
函数格式：
func funcName([parameter1 type])[returnType]{
	// function body
}
*/

// 无参数，无返回值
func sayHello() {
	fmt.Println("Hello world!")
}

// 有参数和返回值：如果所有的参数都是一样的可以省略类型定义
func sum(a, b, c int) int {
	return a + b + c
}

// 返回多个结果值，可以给返回值参数初始化命名
func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b

	return sum, sub
}

// 可变参数，可以传递多个参数
// funName(params ... type)[return]{}
func testVariableParameter(params ...int) int {
	sum := 0
	for i := 0; i < len(params); i++ {
		sum += params[i]
	}

	return sum
}

// defer 语句：定义的语句不执行，只有在函数返回之前执行defer的代码
// 用法1：释放资源，数据库连接句柄，文件句柄等等
// 多个defer：先进后出，相当于栈
func testDefer() {
	//defer fmt.Println("This is first defer code!")
	//defer fmt.Println("This is second defer code!")
	//defer fmt.Println("This is third defer code!")
	for i := 0; i <= 5; i++ {
		defer fmt.Printf("i=%d \n", i)
	}

	fmt.Println("hello")
	fmt.Println("Golang")
}

// defer 使用的变量或者常量是在defer上文中定义的，不会使用defer之后的变量值
func testDeferVar() {
	i := 0
	defer fmt.Printf("defer i=%d \n", i)

	i = 1000
	fmt.Printf("i=%d \n", i)
}

func main() {
	// 简单函数
	sayHello()

	// 多个参数
	sum := sum(1034, 34324, 12334)
	fmt.Println(sum)

	// 多个返回值
	sum, sub := calc(100, 100)
	fmt.Println(sub)

	// 忽略某个返回值使用  _
	sum, _ = calc(10, 10)
	fmt.Println(sum)

	// 可变参数
	multi := testVariableParameter(120, 100)
	fmt.Println(multi)

	// defer
	testDefer()

	// defer
	testDeferVar()
}

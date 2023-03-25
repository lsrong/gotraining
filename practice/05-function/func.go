package main

import "fmt"

/*
函数声明格式：
	func funcName([parameter1 type...])[return type ...] {
		// function body
	}
*/
/*
func main() {
	// func
	sayName()
	fmt.Printf("sum func: %d \n", sum(1, 2, 3))
	calcA, calcB := calc(1, 2)
	fmt.Printf("calc func: %d, %d \n", calcA, calcB)
	fmt.Printf("multiParam func: %d \n", multiParam(1, 2))
	demoDefer()
	demoDeferVar()
}
*/
// sayName 无参数，无返回值 的函数
func sayName() {
	fmt.Println("My name is ShengrongLiu")
}

// sum 有参数和返回值，如果参数或者返回值的类型都一样就可以省略类型定义，最后在定义类型
func sum(a, b, c int) int {
	return a + b + c
}

// calc 多个返回值
func calc(a, b int) (sum, sub int) {
	return a + b, a - b
}

// multiParam 可变参数，可以传递多个参数
// 格式： funcName(params ...type) [return] {}
func multiParam(parameter ...int) int {
	sum := 0
	for i := 0; i < len(parameter); i++ {
		sum += parameter[i]
	}
	return sum
}

// demoDefer defer 在函数放回之前执行的语句（释放资源，如数据库连接句柄，文件句柄等）
// 多个defer 先进后出的顺序执行
func demoDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("demoDefer: i=%d \n", i)
	}
	fmt.Println("demoDefer: Hello world!")
}

// demoDeferVar 使用的变量或者常量是在defer上文中定义的，不会使用defer之后的变量值
func demoDeferVar() {
	i := 0
	defer fmt.Printf("defer i=%d \n", i) // 这种方式的变量的作用域还在当前行
	defer func(j int) {
		fmt.Printf("Last change i=%d \n", j)
	}(i)

	i = 1000
	fmt.Printf("i=%d \n", i)
}

package main

import (
	"fmt"
	"github.com/learning_golang/step1/05-function/visible"
)

// 1. 变量作用域：全局变量和局部变量
// 全局变量：在程序整个生命周期有效

// 定义一个全局变量
var a int = 52

func testGlobalVar() {
	fmt.Printf("Global var a=%d \n", a)
}

//2.局部变量：函数内部定义，语句块内定义，作用在局部范围内,变量重名的时候优先使用局部变量
func testLocalVar() {
	var b int = 1314
	fmt.Printf("Local var b=%d \n", b)

	// 语句块内部
	if true {
		var c int = 55
		fmt.Printf("Local var in if c=%d \n", c)
	}
	if d := 100; d > 0 {
		fmt.Printf("Local var in if d=%d \n", d)
	} else {
		fmt.Printf("else d=%d \n", d)
	}
	//fmt.Printf("Local var in if d=%d \n", d)
}

// 访问控制
// 函数，变量，常量等等：首字母大写为公用，小写为私有

func main() {
	testGlobalVar()
	testLocalVar()
	fmt.Printf("This is visible function:Add ret is %d \n", visible.Add(visible.A, 100))
	//visible.sub(1, 1)

}

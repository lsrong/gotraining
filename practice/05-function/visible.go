package main

import (
	"fmt"

	"github.com/learning_golang/practice/05-function/visible"
)

// 定义全局变量
var a int = 520

/*
*
// 全局变量：在程序整个生命周期有效
// 局部变量：函数内部定义，语句块内定义，作用在局部范围内,变量重名的时候优先使用局部变量
*/
/*
func main() {
	demoGlobalVar()
	demoLocalVar()
	demoVisible()
}
*/

// demoGlobalVar 全局变量：在程序整个生命周期有效
func demoGlobalVar() {
	fmt.Printf("Global var a =%d \n", a)
}

// demoLocalVar 函数内部定义，语句块内定义，作用在局部范围内,变量重名的时候优先使用局部变量
func demoLocalVar() {
	var b int = 1314
	fmt.Printf("Local var b = %d \n", b)
	// 语句内部
	if true {
		var c int = 55
		fmt.Printf("Local var in if c=%d \n", c)
	}

	if d := 100; d > 0 {
		fmt.Printf("Local var in if d =%d \n", d)
	} else {
		fmt.Printf("else d= %d \n", d)
	}
}

// demoVisible 访问控制：函数，变量，常量等等：首字母大写为公用，小写为私有
func demoVisible() {
	fmt.Printf("This is visible function:Add ret is %d \n", visible.Add(visible.A, 100))
}

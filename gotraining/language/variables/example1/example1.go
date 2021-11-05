package main

import "fmt"

//Sample program to show how to declare variables
func main() {

	/**
	Zero Values(Type Initialized Value): "零值"

	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil

	*/

	// Declare variables that are set to their zero value 声明变量,默认会赋零值
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v] \n", a, a)
	fmt.Printf("var b string \t %T [%v] \n", b, b)
	fmt.Printf("var c float64 \t %T [%v] \n", c, c)
	fmt.Printf("var d bool \t %T [%v] \n\n", d, d)

	// Declare variables and initialize 声明并初始化变量
	// Using the short variable declaration operator 使用短赋值操作符号 :=
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa :=10 \t %T [%v] \n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v] \n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v] \n", cc, cc)
	fmt.Printf("dd := true \t %T [%v] \n\n", dd, dd)

	// Specify type and preform a conversion. 指定类型转换操作.
	aaa := int32(10)

	fmt.Printf("aaa := int62(10) %T [%v] \n", aaa, aaa)

}

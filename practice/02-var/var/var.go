package main

import "fmt"

// 变量
func testVar() {
	//  三种声明方式
	var a int
	var b = 10
	c := "string"
	a = 1000
	fmt.Printf("三种声明方式: a = %d, b = %d, c = %s \n", a, b, c)

	// 多变量声明
	var d, e string
	d = "hello"
	e = "second"
	fmt.Printf("多变量声明：d=%s, e=%s \n", d, e)

	//  声明并赋值
	f, g := 1, 2
	var (
		h string = "h string"
		i int    = 520
	)
	fmt.Printf("声明并赋值：f= %d, g=%d, h=%s, i=%d\n", f, g, h, i)

	// 变量互换
	m, n := 1, 2
	fmt.Printf("变量互换： m=%d, n=%d", m, n)
	m, n = n, m
	fmt.Printf(", m=%d, n=%d \n", m, n)

	// _ 丢弃变量
	_ = n
	fmt.Printf("n=%d \n\n", n)
}

// testConst 常量定义
func testConst() {
	// 多常量定义
	const (
		pi     = 3.14159
		prefix = "Go_"
	)
	fmt.Printf("%f, %s \n", pi, prefix)

	// 关键字iota
	const (
		itA = iota // 0
		itB        // 1
		itC        // 2
	)
	fmt.Printf("%d, %d, %d \n", itA, itB, itC)

	// iota 在同一行的时候值都是一样的
	const (
		itG           = iota             // 0
		itH, itI, itJ = iota, iota, iota //1,1,1
		itK           = iota             // 2
	)
	fmt.Printf("%d, %d, %d, %d, %d \n", itG, itH, itI, itJ, itK)

	const (
		itL = 1 << iota // 1
		itM             // 2
		itN             // 4
	)
	fmt.Printf("%d, %d, %d \n", itL, itM, itN)
}

type User struct {
	Name string
	Age  int
}

// testStruct 结构体定义
func testStruct() {
	user := User{
		Name: "Hello",
		Age:  30,
	}
	fmt.Printf("输出结构体变量：%v \n", user)
}

func main() {
	testVar()
	
	testConst()

	testStruct()
}

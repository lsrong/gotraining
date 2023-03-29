package main

import "fmt"

func initPoint() {
	// 地址：&
	var a int
	a = 100

	// 引用定义
	var b *int
	b = &a

	fmt.Printf("a=%d \n", a)
	fmt.Printf("b=%d \n", *b)
}

// 指针类型修改
func modify(p *int) {
	*p = 10000
}
func testModify() {
	var a = new(int)
	*a = 10
	fmt.Printf("before modify a=%d \n", *a)
	modify(a)
	fmt.Printf("after modify a=%d \n", *a)
}

// 数组类型指针
func modifyArrayPoint(arr *[3]int) {
	(*arr)[0] = 100
	(*arr)[1] = 200
	(*arr)[2] = 300
}
func testModifyArrayPoint() {
	var a = [3]int{1, 2, 3}
	fmt.Print("before modify array")
	fmt.Println(a)

	modifyArrayPoint(&a)
	fmt.Print("after modify array")
	fmt.Println(a)
}

/**
new与make的区别
二者都是用来做内存分配的。
make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
func testNew() {
	//var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋
	var a *int
	var b int
	a = &b
	b = 100
	fmt.Printf("a=%d \n", *a)
	fmt.Printf("b=%d \n", b)

	// 使用内置的new函数对a进行初始化之后就可以正常对其赋值
	var c = new(int) // 分配指针
	*c = 100
	fmt.Printf("c=%d \n", *c)
}

// 如果指针变量存放同一个内存指针的时候，修改的是同一块内存空间的值
func testManyPointParam() {
	var a int
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("a=%d \n", a)
	fmt.Printf("b=%d \n", *b)
	fmt.Printf("c=%d \n", *c)
}

/*
func main() {
	initPoint()
	testModify()
	testModifyArrayPoint()
	testNew()
	testManyPointParam()
}
*/

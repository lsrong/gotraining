package main

import "fmt"

func main() {
	demoInitPoint()
	demoNew()
}

// demoInitPoint 指针类型的增删改查
func demoInitPoint() {
	// 指针类型 & 取指符号
	var (
		a int
		b *int
	)
	a, b = 100, &a
	fmt.Printf("demo point a=%d\n", a)
	// 读取指针
	fmt.Printf("demo point b=%d\n", *b)

	// 指针类型的修改
	modify := func(p *int) {
		*p = 10000
	}
	// new 初始化一个指针类型的变量
	var c = new(int)
	*c = 10
	fmt.Printf("before modify a=%d \n", *c)
	modify(c)
	fmt.Printf("after modify a=%d \n", *c)

	modifyArrayPoint := func(arr *[3]int) {
		(*arr)[0] = 100
		(*arr)[1] = 200
		(*arr)[2] = 300
	}
	var d = [3]int{1, 2, 3}
	fmt.Printf("before modify array: d=%v \n", d)
	modifyArrayPoint(&d)
	fmt.Printf("after modify array:d=%v \n", d)
}

// demoNew 创建指针的方式：new, make
// new与make的区别
// 二者都是用来做内存分配的。
// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func demoNew() {
	var a *int // 只是声明了一个指针类型但是没有初始化，不能给 a 赋值，（未初始化）
	// a = 100	// 报错：cannot use 100 (untyped int constant) as *int value in assignmentgo
	var b int
	a = &b // 指针作为引用类型需要 初始化才会拥有内存空间，才可以给它赋值
	b = 100
	fmt.Printf("a=%d \n", *a)
	fmt.Printf("b=%d \n", b)

	var c = new(int) // new()函数对a进行初始化之后就可以正常赋值了
	*c = 100
	fmt.Printf("c=%d \n", *c)
}

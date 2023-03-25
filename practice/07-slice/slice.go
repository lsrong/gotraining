package main

import "fmt"

func main() {
	demoSlice()
	demoSliceFunc()
}

// demoSlice 切片示例， 基于数组类型的一层封装，自动扩容
func demoSlice() {
	// 切片的零值为 nil
	var a []int
	if a == nil {
		fmt.Println("切片的零值为nil")
	}

	// 初始化，a[start:end], 下标 start 到  end-1 的映射 切片[a, b)
	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bSlice := b[4:6]
	fmt.Printf("slice bSlice= %v, type of bSlice is %T\n", bSlice, bSlice)

	// 定义并初始化
	c := []int{1, 2, 3}
	fmt.Printf("slice c = %v, type of c is %T \n", c, c)
	fmt.Println()
	/**
	数组切⽚的基本操作:
		a) arr[start:end]：包括start到end-1(包括end-1)之间的所有元素
		b) arr[start:]：包括start到arr最后⼀个元素(包括最后⼀个元素)之间的所有元素
		c) arr[:end]：包括0到end-1(包括end-1）之间的所有元素
		d) arr[:]：包括整个数组的所有元素
	*/
	// arr[start:end]
	d := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	e := d[4:6]
	fmt.Printf("arr[start:end]: %v \n", e)
	// arr[start:]
	f := d[1:]
	fmt.Printf("arr[start:]: %v \n", f)
	// arr[:end]
	g := d[:6]
	fmt.Printf("arr[:end]: %v \n", g)
	// arr[:]
	h := d[:]
	fmt.Printf("arr[:] = %v, type of c is %T \n", h, h)

	// 遍历
	for i, v := range h {
		fmt.Printf("h[%d] = %d \n", i, v)
	}

	// 切片为引用类型
	i := [...]int{1, 2, 3}
	j := i[:]
	j[0] = 100
	fmt.Printf("切片为引用类型:i = %v \n", i)
}

// demoSliceFunc 切片相关的标准函数
func demoSliceFunc() {
	// make(type , len, cap) 定义和初始化切片的底层数组的长度容量, 默认值为对应类型的零值
	a := make([]int, 3, 5)
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Printf("make a = %v \n", a)

	// append 追加元素到切片
	a = append(a, 4)
	fmt.Printf("append a = %v \n", a)

	// 自动扩容
	for i := 0; i < 10; i++ {
		a = append(a, i+10)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	fmt.Println(a)

	//cap 切片的容量， 赋值一份切片
	b := a[:cap(a)]
	fmt.Println(b)
}

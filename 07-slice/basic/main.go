package main

import "fmt"

// 切片：基于数组类型的一层封装，自动扩容
func initSlice() {
	// 空切片
	var a []int
	fmt.Println(a)
	// nil 零值

	// 初始化，a[start:end],从 start 到 end-1,（下标）
	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bSlice := b[4:6]
	fmt.Println(bSlice)
	fmt.Printf("slice bSlice:%v type of bSlice:%T\n", a, a)

	c := []int{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("slice c:%v type of c:%T\n", c, c)
}

func sliceOperation() {
	/**
	数组切⽚片基本操作:
		a) arr[start:end]：包括start到end-1(包括end-1)之间的所有元素
		b) arr[start:]：包括start到arr最后⼀一个元素(包括最后⼀一个元素)之间的所有元素
		c) arr[:end]：包括0到end-1(包括end-1）之间的所有元素
		d) arr[:]：包括整个数组的所有元素
	*/
	// arr[start:end]
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[4:6]
	fmt.Println(b)

	// arr[start:]
	c := a[1:]
	fmt.Println(c)

	// arr[:end]
	d := a[:5]
	fmt.Println(d)

	// arr[:]
	e := a[:]
	fmt.Println(e)

	// 遍历
	for i, v := range b {
		fmt.Printf("b[%d] = %d \n", i, v)
	}
}

// 切片为引用类型
func sliceType() {
	a := [...]int{1, 2, 3}
	b := a[:]
	b[0] = 100
	fmt.Println(a)
}

func main() {
	initSlice()

	sliceOperation()

	sliceType()
}

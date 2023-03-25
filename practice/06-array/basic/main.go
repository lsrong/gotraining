package main

import "fmt"

// 定义数组类型：长度 + 类型
// 下标从0开始
func testArray() {
	var a [10]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [5]int{6, 7, 8, 9, 10}
	fmt.Println(c)

	// [...]:长度为定义的最长下标+1
	d := [...]int{0, 1, 2, 3, 4}
	fmt.Println(d)

	// 初始化指定下标， index:value
	e := [...]int{4: 10, 6: 20}
	fmt.Println(e)
	fmt.Println(len(e))

	var f [10]int
	// f = e 长度+值类型 组成数据类型
	f = a
	fmt.Println(f)
}

// 数组遍历：1. for 2. for range
func testIterate() {
	a := [5]int{2: 2, 3: 10}
	// for
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d \n", i, a[i])
	}

	fmt.Println("for i,v :=range")
	// for := range
	for i, v := range a {
		fmt.Printf("a[%d]=%d \n", i, v)
	}

	// _ , v :=range
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}
}

// ç
func testMultiArrays() {
	a := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}

	fmt.Println(a)
	fmt.Println(len(a))

	// 遍历
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d \n", i, j, a[i][j])
		}
	}

	fmt.Println("for i,v :=range")

	for i, v1 := range a {
		for j, v2 := range v1 {
			fmt.Printf("a[%d][%d] = %d \n", i, j, v2)
		}
	}

}

func main() {
	testArray()

	testIterate()

	testMultiArrays()
}

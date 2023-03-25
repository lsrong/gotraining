package main

import (
	"fmt"
	"sort"
)

/**
func main() {
	demoAppend()
	demoSortArray()
}
*/

// demoAppend 追加切片元素
func demoAppend() {
	var a = make([]string, 2, 5)
	fmt.Printf("make slice a = %v \n", a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%d", i))
	}
	fmt.Printf("append slice a = %v \n", a)
}

// demoSortArray 使用golang标准包“sort”对数组排序
func demoSortArray() {
	a := []int{0, 123, 4325, 54, 5235}
	sort.Ints(a)
	fmt.Printf("sort.Ints a=%v \n", a)

	b := []string{"c", "e", "d", "a"}
	sort.Strings(b)
	fmt.Printf("sort.Stings b=%v \n", b)

	c := []float64{0.89, 1.1, 45.00, 2.2, 3.14159}
	sort.Float64s(c)
	fmt.Printf("sort.Float64s c=%v \n", c)
}

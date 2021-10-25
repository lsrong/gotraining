package main

import (
	"fmt"
	"sort"
)

// 追加数组切片
func makeAppend() {
	var a = make([]string, 2, 5)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%d", i))
	}
	fmt.Println(a)
}

// 使用golang标准包“sort”对数组排序
func sortArray() {
	a := []int{0, 2, 4, 1, 7, 5, 2}
	fmt.Println(a)
	sort.Ints(a)

	b := []string{"c", "e", "d", "a"}
	sort.Strings(b)
	fmt.Println(b)

	c := []float64{0.8, 0.21, 1.54, 0.1}
	sort.Float64s(c)
	fmt.Println(c)
}

func main() {
	makeAppend()

	sortArray()
}

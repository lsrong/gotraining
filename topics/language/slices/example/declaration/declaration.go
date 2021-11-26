package main

import "fmt"

// Sample program to show how the capacity of the slice
// is not available for use.

// 显示切片容量如何不可用的示例程序。切片的长度不能超过容量

func main() {
	// make 第三个参数没传默认cap = len
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// 元素值不允许大于容量数量,除非使用append函数自动扩容追加元素
	fruits[5] = "Out of capacity"
	// panic: runtime error: index out of range [5] with length 5

	fmt.Println(fruits)
}

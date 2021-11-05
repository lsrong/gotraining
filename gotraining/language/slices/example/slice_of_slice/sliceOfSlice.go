package main

import "fmt"

// Sample program to show how to takes slices of slices to create different
// views of and make changes to the underlying array.
// 示例程序显示如何获取切片以创建不同的视图并对底层数组进行更改。

// 从切片中生成新切片: slice := slice[n:m], 半开区间[) 获取元素, 生成的切片长度为 m-n, 底层数据不变.

func main() {
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)]
	// 取一个切片, 获取下标为2,3的元素
	// 获取参数: [starting_index : (starting_index + length)] = [starting_index: end_index],
	// 所以新切片的长度为: length = end_index - string_index
	slice2 := slice1[2:4]
	inspectSlice(slice2)

	fmt.Println("*************************")

	// slice1, slice2都是对同一个底层数组的引用,所以如果改变其中一个切片的元素值,另外一个切片对应下标的元素值也会改变
	slice2[0] = "CHANGED"
	inspectSlice(slice1) // [2] 0xc000106020 CHANGED
	inspectSlice(slice2) // [0] 0xc000106020 CHANGED

	fmt.Println("*************************")
	// 如果要避免引用同一个底层数组,可以重新创建新的空切片并将通过内建函数copy复制到新切片中

	// Make a new slice big enough to hold elements of slice 1 and copy the
	// values over using the builtin copy function.
	// 制作一个足够大的新切片以容纳切片 1 的元素，
	// 并使用内置复制函数复制值。
	slice3 := make([]string, len(slice1))
	copy(slice3, slice1)
	// slice3会分配新的内存空间.
	inspectSlice(slice3)
}

func inspectSlice(slice []string) {
	fmt.Printf("Len[%d] Cap[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}

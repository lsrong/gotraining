package main

import "fmt"

// Sample program to show how to declare and iterate over
// arrays of different types.

//演示定义与遍历数组, 长度(正数) + 类型都是数组类型的一部分.

func main() {
	// 定义5个字符串长度的数组,初始化为空.
	var colors [5]string
	colors[0] = "red"
	colors[1] = "yellow"
	colors[2] = "blue"
	colors[3] = "black"
	colors[4] = "white"

	// for range 遍历
	for i, color := range colors {
		fmt.Println(i, color)
	}

	// 定义并初始化数组
	numbers := [4]int{10, 20, 30, 40}

	// for len 遍历
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// 不同的定义数组方式
	displayArr()
}

func displayArr() {
	// ... 根据初始化给定的元素设置数组,未初始化会自动设置为零值
	b := [...]int{1, 2, 3, 4} //透過初始化給的元素數量來給定長度
	fmt.Println(b, len(b))    // [1 2 3 4] 4

	v := [...]int{1: 2, 3: 4} //透過索引初始化元素, 沒被初始化的就是該類型的預設值
	fmt.Println(v, len(v))    // [0 2 0 4] 4
}

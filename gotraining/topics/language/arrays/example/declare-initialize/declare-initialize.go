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

}

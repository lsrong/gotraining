package main

import "fmt"

// Sample program to show how the for range has both value and pointer semantics.

// 演示range分别遍历值和指针数组.

// 数组属于值类型,一定义就内存分配就已经固定.

func main() {
	colors := [5]string{"Red", "Blue", "Pink", "Black", "White"}
	fmt.Printf("Bfr[%s] : ", colors[1])
	// Using the pointer semantic form of the for range.
	// 使用 for 范围的指针语义形式, 修改数组原值, array[i] = value
	for i := range colors {
		colors[1] = "Yellow"
		if i == 1 {
			fmt.Printf("Aft[%s]\n", colors[1])
		}
	}

	// Using the value semantic form of the for range.
	// 使用 for 范围的值语义形式。
	colors = [5]string{"Red", "Blue", "Pink", "Black", "White"}
	fmt.Printf("Bfr[%s] : ", colors[1])
	for i, v := range colors {
		colors[1] = "Yellow"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}

	// Using the value semantic form of the for range but with pointer semantic access. DON'T DO THIS.
	// 使用 for 范围的值语义形式但具有指针语义访问。不要这样做。
	colors = [5]string{"Red", "Blue", "Pink", "Black", "White"}
	fmt.Printf("Bfr[%s] : ", colors[1])
	for i, v := range &colors {
		colors[1] = "Yellow"
		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}

}

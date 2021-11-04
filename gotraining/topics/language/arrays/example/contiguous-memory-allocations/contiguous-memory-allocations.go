package main

import "fmt"

// Sample program to show how the behavior of the for range and
// how memory for an array is contiguous.
// 示例程序显示 for 范围的行为以及数组的内存如何连续。

// Memory is allocated as a contiguous block.
// 数组类型的内存被分配为一个连续的块。

func main() {
	colors := [5]string{"Red", "Blue", "Pink", "Black", "White"}

	// range 的 i, v为临时变量, 分配
	for i, v := range colors {
		fmt.Printf("Index[%d %p]\t Value[%s %p]\t Address[%p]\n", i, &i, v, &v, &colors[i])
	}
	/**
	 * output: Address is contiguously
	Index[0 0xc0000200f0]    Value[Red 0xc000010240]         Address[0xc00006a050]
	Index[1 0xc0000200f0]    Value[Blue 0xc000010240]        Address[0xc00006a060]
	Index[2 0xc0000200f0]    Value[Pink 0xc000010240]        Address[0xc00006a070]
	Index[3 0xc0000200f0]    Value[Black 0xc000010240]       Address[0xc00006a080]
	Index[4 0xc0000200f0]    Value[White 0xc000010240]       Address[0xc00006a090]
	*/
}

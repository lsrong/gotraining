package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	Value1 int
	Value2 string
}

func main() {
	// 初始化'值'类型
	myValue := MyType{10, "Bill"}
	pointer := unsafe.Pointer(&myValue)
	fmt.Printf("1. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)

	// 传递值副本给函数
	changeMyValue(myValue)

	// 调用完毕堆栈会弹出myValue副本.
	fmt.Printf("3. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

// 	myValue实际为参数副本, 堆栈中会开辟新内存空间拷贝myValue,作为参数传递给函数调用
func changeMyValue(myValue MyType) {
	// 改变myValue结构体值,
	myValue.Value1 = 20
	myValue.Value2 = "Jill"

	// Create a pointer to the memory for myValue
	pointer := unsafe.Pointer(&myValue)

	// Display the address and values
	fmt.Printf("2. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

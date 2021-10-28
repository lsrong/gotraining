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
	// 初始化引用类型
	myValue := &MyType{10, "Bill"}
	pointer := unsafe.Pointer(myValue)
	fmt.Printf("1. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)

	// 传递指针给函数
	changeMyValue(myValue)

	// 函数changMyValue修改的数据都是同一个数据块
	fmt.Printf("3. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

// 	myValue参数的'值'为MyType指针,堆栈中开辟保存的为指针,此时myValue可以表示为指针变量,`值为指针的变量`
func changeMyValue(myValue *MyType) {
	// 改变myValue结构体值,
	myValue.Value1 = 20
	myValue.Value2 = "Jill"

	// Create a pointer to the memory for myValue
	pointer := unsafe.Pointer(myValue)

	// Display the address and values
	fmt.Printf("2. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

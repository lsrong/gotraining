package main

import (
	"fmt"
	"unsafe"
)

func main() {
	myMap := make(map[string]string)
	myMap["Bill"] = "Jill"

	pointer := unsafe.Pointer(&myMap)
	fmt.Printf("1. Addr: %v Value : %s\n", pointer, myMap["Bill"])

	changeMyMap(myMap)
	fmt.Printf("3. Addr: %v Value : %s\n", pointer, myMap["Bill"])

	changeMyMapAddr(&myMap)
	fmt.Printf("5. Addr: %v Value : %s\n", pointer, myMap["Bill"])

}

// 传递为map数据结构,由于map是引用类型,只会复制数据结构,而不会复制底层数据,因此myMap参数具有与上次相同映射的变量,栈地址不一样而已
func changeMyMap(myMap map[string]string) {
	myMap["Bill"] = "Joan"
	pointer := unsafe.Pointer(&myMap)

	fmt.Printf("2. Addr: %v Value : %s\n", pointer, myMap["Bill"])
}

// 传递为map的引用地址,具有数据结构不会复制,而是传入指针地址
func changeMyMapAddr(myMapPointer *map[string]string) {
	(*myMapPointer)["Bill"] = "Jenny"
	pointer := unsafe.Pointer(myMapPointer)

	fmt.Printf("4. Addr: %v Value : %s\n", pointer, (*myMapPointer)["Bill"])
}

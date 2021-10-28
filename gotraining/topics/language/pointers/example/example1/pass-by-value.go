package main

import "unsafe"

// Sample program to show the basic concept of pass by value.
// 按值传递简单示例

func main() {
	count := 10
	println("1. count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// pass by value, 按值传递
	incrementVar(count)
	println("3. count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// pass by pointer 按值指针传递
	incrementPointer(&count)
	println("5. count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

}

func incrementVar(incr int) {
	incr++
	println("2. count:\tValue Of[", incr, "]\tAddr Of[", &incr, "]")
}

func incrementPointer(incr *int) {
	*incr++
	pointer := unsafe.Pointer(incr)
	println("4. count:\tValue Of[", incr, "]\tAddr Of incr[", &incr, "] Addr of count[", pointer, "]")
}

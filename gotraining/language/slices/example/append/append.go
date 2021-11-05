package main

import "fmt"

// Sample program to show how to grow a slice using the built-in function append
// and how append grows the capacity of the underlying array.
// 示例程序显示如何使用内置函数 append 增加切片以及 append 如何增加底层数组的容量。

// 内建函数append对底层数据实现自动扩容增

// 每次扩容都是重新拷贝底层数组到更大长度的数组里面.

func main() {
	var data []string
	lastCap := cap(data)

	for i := 1; i <= 1e5; i++ {
		value := fmt.Sprintf("Rec: %d", i)
		data = append(data, value)

		if lastCap != cap(data) {
			// 容量变化的百分比.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// 变化的容量
			lastCap = cap(data)

			fmt.Printf("Addr[%p] \t Index[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0], i, lastCap, capChg)
		}
	}

	var demo []byte
	demo1 := append(demo, []byte("hello")...)
	fmt.Println("aft append", demo1)

	demo2 := AppendByte(demo, []byte("hello")...)
	fmt.Println("aft append", demo2)
}

// AppendByte 模拟append函数
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)

	return slice
}

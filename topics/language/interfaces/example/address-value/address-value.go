package main

// Sample program to show how you can't always get the address of a value.
// 示例程序显示如何无法始终获取值的地址。

// 强制装换T(), 函数返回F(),的值是不可寻址的,不能直接调用指针方法.

import "fmt"

type duration int64

func (d *duration) notify() {
	fmt.Println("Sending notification in", *d)
}

func newDuration(n int64) duration {
	return duration(n)
}

func main() {
	duration(100).notify()
	// ./address-value.go:12:15: cannot call pointer method on duration(100)
	//./address-value.go:12:15: cannot take the address of duration(100)

	d := duration(100)
	d.notify()

	newDuration(100).notify()
	//./address-value.go:26:18: cannot call pointer method on newDuration(100)
	//./address-value.go:26:18: cannot take the address of newDuration(100)

	d = newDuration(100)
	d.notify()
}

package main

import "fmt"

// Sample program to show how iota works.
// iota 常量的使用规则

func main() {
	const (
		A1 = iota // 0 iota从0开始
		B1 = iota // 1 递增1
		C1 = iota // 2 递增1
	)
	fmt.Println("1:", A1, B1, C1)

	// 省略iota
	const (
		A2 = iota // 0
		B2        // 1
		C2        // 2
	)
	fmt.Println("2:", A2, B2, C2)

	// 设置开始数值
	const (
		A3 = iota + 1 // 1: Start at 0 +1
		B3            // 2
		C3            // 3
	)
	fmt.Println("3:", A3, B3, C3)

	// 跳过某个数
	const (
		A4 = iota // 0 Start at 0
		_         // ignore 1
		B4        // 2
		_         // ignore 3
		C4        // 4
	)
	fmt.Println("4:", A4, B4, C4)

	// 左移操作
	const (
		Ldate           = 1 << iota // 1 向左移动 0(iota) 位, 0000,0001
		Ltime                       // 2 向左移动 1(iota) 位, 0000,0010
		Lmicroseseconds             // 4 向左移动 2(iota) 位, 0000,0100
		Llongfile                   // 8 向左移动 3(iota) 位, 0000,1000
		Lshortfile                  // 16 向左移动 4(iota) 位, 0001,0000
		LUTC                        // 32 向左移动 5(iota) 位, 0010,0000
	)

	fmt.Println("log:", Ldate, Ltime, Lmicroseseconds, Llongfile, Lshortfile, LUTC)

}

package main

import "fmt"

// Sample program to show how arrays of different sizes are
// not of the same type.

// 长度和元素类型共同组成数组类型.
// 不同大小的相同类型不属于同一类型数组。

func main() {
	var five [5]int

	four := [4]int{10, 20, 30, 40}

	// 长度不同类型也不同
	//five = four // ./different-type.go:17:7: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(five)
	fmt.Println(four)
}

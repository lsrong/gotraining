package main

import "fmt"

// Sample program to show how constants do have a parallel type system.
// 演示常量如何具有并行类型系统
// parallel type : https://zh.wikipedia.org/wiki/%E5%B9%B6%E8%A1%8C%E8%AE%A1%E7%AE%97

const (
	// 比int64大的数字
	bigger = 87989787564654645646548745259456

	// 如果指定类型则常量有对应类型精度限制
	// ./parallel-type.go:14:2: constant 879897875646546456465487452594565645646 overflows int64
	//biggerInt int64 = 879897875646546456465487452594565645646
)

func main() {
	fmt.Println("Will Compile")
}

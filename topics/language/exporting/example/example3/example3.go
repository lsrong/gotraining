package main

import (
	"fmt"
	"github.com/learning_golang/topics/language/exporting/example/example3/counters"
)

// Sample program to show how the program can access a value
// of an unexported identifier from another package.
// 示例程序显示程序如何访问来自另一个包的未导出标识符的值。

// 使用辅助函数间接使用未导出的类型.

func main() {
	// 通过New辅助方法创建未导出的类型.
	counter := counters.New(10)

	fmt.Printf("Counters: %d \n", counter)
}

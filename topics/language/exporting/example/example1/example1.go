package main

// Sample program to show how to access an exported identifier.
// 显示如何访问导出标识符(package)的示例程序。

import (
	"fmt"
	"github.com/learning_golang/topics/language/exporting/example/example1/counters"
)

func main() {
	// 初始化定义counters.AlertCounters类型
	counter := counters.AlertCounter(100)

	fmt.Printf("Counters: %d", counter)
}

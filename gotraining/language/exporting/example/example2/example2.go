package main

// Sample program to show how the program can't access an
// unexported identifier from another package.
// 显示程序如何无法从另一个包访问未导出的标识符的示例程序。
import (
	"fmt"
	"github.com/learning_golang/gotraining/language/exporting/example/example2/counters"
)

func main() {
	// 不能直接使用未导出的类型
	counter := counters.alertCounters(10)

	// ./example2.go:9:13: cannot refer to unexported name counters.alertCounters

	fmt.Printf("Counter: %d \n", counter)
}

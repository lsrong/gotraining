package main

import (
	"fmt"
	"github.com/learning_golang/gotraining/language/exporting/example/example5/users"
)

// Sample program to show how to create values from exported types with
// embedded unexported types.
// 示例程序显示如何从带有嵌入的未导出类型的导出类型创建值。

// 先初始化后赋值.

func main() {
	u := users.Manager{
		Title: "Admin",
	}

	// 设置user已导出的属性.
	u.Name = "Li"
	u.ID = 1

	fmt.Printf("User: %#v\n", u)
}

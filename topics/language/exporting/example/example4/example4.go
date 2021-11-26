package main

import (
	"fmt"
	"github.com/learning_golang/topics/language/exporting/example/example4/users"
)

// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.

// 示例程序显示如何无法直接访问来自导出结构类型的未导出字段。

func main() {
	u := users.User{
		ID:   1,
		Name: "li",

		password: "123456",
	}

	// ./example4.go:18:3: cannot refer to unexported field 'password' in struct literal of type users.User

	fmt.Printf("User: %#v\n", u)
}

package main

import "fmt"

// Sample program to show how method sets can affect behavior.

// 显示方法集如何影响行为的示例程序。

// 行为会随着方法集的不同而变化.

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	return fmt.Sprintf("My name is %s and email is %s", u.name, u.email)
}

func main() {
	u := user{
		name:  "Li",
		email: "Li@email.com",
	}

	// 会打印结构体
	fmt.Println(u)

	// 会打印String(), 指针方法集实现了接口fmt.Stringer,执行String()行为
	fmt.Println(&u)
}

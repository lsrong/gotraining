package main

import "fmt"

// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
// 示例程序展示了如何将一个类型嵌入到另一个类型中以及内部和外部类型之间的关系。

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin 嵌入user 类型.
type admin struct {
	user  // Embedded Type
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "Li",
			email: "Li@email.com",
		},
		level: "super",
	}

	// 调用admin嵌入类型user的方法
	ad.user.notify()

	// 嵌入类型的方法可被提升提供上层结构体调用.
	ad.notify()
}

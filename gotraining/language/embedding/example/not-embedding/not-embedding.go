package main

import "fmt"

// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
//示例程序展示了我们所做的不是嵌入类型，而是仅使用类型作为字段

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin 结构体中person不是类型嵌入,而是属性字段, 属性值为user, user 方法不会升级为admin的方法.
type admin struct {
	person user
	level  string
}

func main() {
	ad := admin{
		person: user{
			name:  "Li",
			email: "Li@email.com",
		},
		level: "super",
	}

	ad.person.notify()

	//ad.notify()
	// ./not-embedding.go:35:4: ad.notify undefined (type admin has no field or method notify)
}

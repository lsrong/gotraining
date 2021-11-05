package main

import "fmt"

// Sample program to show the basic concept of using a pointer
// to share data.

// 通过指针共享数据

type user struct {
	name   string
	email  string
	logins int
}

func main() {
	bill := user{
		name:  "Bill",
		email: "Bill@gmail.com",
	}

	// 指针作为值传递
	display(&bill)

	// 传递结构体中logins字段的指针作为值传递
	increment(&bill.logins)

	display(&bill)
}

func display(user *user) {
	fmt.Printf("%p \t %+v \n", user, *user)
	fmt.Printf("Name:%s, Email:%s, Logins:%d \n", user.name, user.email, user.logins)
}

func increment(logins *int) {
	// &logins 表示值指针的指针地址.
	// logins 表示值的指针地址.
	// *logins 表示值指针指向的数据.
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n", &logins, logins, *logins)
}

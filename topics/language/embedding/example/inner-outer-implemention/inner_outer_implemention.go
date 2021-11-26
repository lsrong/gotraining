package main

// Sample program to show what happens when the outer and inner
// type implement the same interface.
// 示例程序显示当外部和内部类型实现相同的接口时会发生什么。

// - 如果外部类型和内部类型同时实现了相同的方法,外部调用首先调用外部类型的方法.
// - 如果外部类型和内部类型同时实现了某个接口, 调用接口变量时为外部类型的方法.
// - 优先级: 外部类型方法 > ...> 内部类型方法

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email To %s<%s>\n", a.name, a.email)
}
func main() {
	ad := admin{
		user: user{
			name:  "Li",
			email: "Li@email.com",
		},
		level: "super",
	}
	sendNotification(&ad)

	// 可以通过内部属性调用嵌入类型的方法,属性名 == 类型名
	ad.user.notify()

	// 	内部嵌入类型的方法不会提升.
	ad.notify()

}

func sendNotification(n notifier) {
	n.notify()
}

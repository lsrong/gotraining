package main

import "fmt"

// Sample program to show how embedded types work with interfaces.
// 显示嵌入类型如何与接口一起工作的示例程序。
// - 接口的嵌入内部类型的实现被“提升”到外部类型。
// - 内部嵌入类型实现某个接口则外部类型就实现了该接口

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

// admin 内嵌类型user实现接口notifier,则admin也实现该接口
type admin struct {
	user
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

	// sendNotification(ad)
	// ./embedding-interface.go:37:18: cannot use ad (type admin) as type notifier in argument to sendNotification:
	//	admin does not implement notifier (notify method has pointer receiver)

	// 接口的嵌入内部类型的实现被“提升”到外部类型。
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}

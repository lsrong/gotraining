package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User email to %s<%s>", u.name, u.email)
}

func main() {
	u := user{name: "LI", email: "LI@email.com"}

	// ./method-set.go:21:18: cannot use u (type user) as type notifier in argument to sendNotification:
	//	user does not implement notifier (notify method has pointer receiver)
	// user 类型的值不实现接口，因为指针接收器不属于值的方法集。
	//sendNotification(u)

	// user 类型的指针实现了接口.
	sendNotification(&u)

	// 值方法集合: 值接受者实现的方法
	// 指针方法集合: 值+指针接受者实现的方法

}

// sendNotification 接受实现通知程序接口的值并发送通知。
// 定义数据抽象行为
func sendNotification(n notifier) {
	n.notify()
}

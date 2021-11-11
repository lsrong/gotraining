package main

import "fmt"

// Sample program to show how to declare and use function types.
// 演示如何声明和使用函数类型的示例程序。

func event(message string) {
	fmt.Println(message)
}

type data struct {
	name string
	age  int
}

func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

func fireEvent1(f func(string)) {
	f("anonymous")
}

// handler 定义代表处理时间的函数类型
type handler func(string)

// fireEvent2 使用函数类型传参
func fireEvent2(h handler) {
	h("handler")
}

func main() {
	d := data{
		name: "Li",
	}

	// 使用接受任何具有正确签名的函数或方法的 fireEvent1 处理程序。
	// 方法和函数传参时候如果签名(名称)一致都可以作为匿名函数参数传参.
	fireEvent1(event)
	fireEvent1(d.event)

	// 使用接受任何类型为 `handler` 的函数或方法
	// 或任何具有正确签名的文字函数或方法的 fireEvent2 处理程序。
	fireEvent2(event)
	fireEvent2(d.event)

	// 为基于全局和方法的事件函数声明一个处理程序类型的变量。
	h1 := handler(event)
	h2 := handler(d.event)

	// fireEvent1 和 fireEvent2都可以接受签名相同的handler类型变量或者匿名函数.
	fireEvent1(h1)
	fireEvent1(h2)

	fireEvent2(h1)
	fireEvent2(h2)
}

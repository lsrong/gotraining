package main

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
// 使用名为 speak 的方法声明一个名为 Speaker 的接口。
// 声明一个名为 english 的结构体，代表一个会说英语的人，并为一个会说中文的人声明一个名为 chinese 的结构体。
// 使用值接收器和这些文字字符串“Hello World”和“你好世界”为每个结构实现扬声器接口。
// 声明一个speaker 类型的变量并分配一个english 类型值的地址并调用该方法。
// 对 chinese 类型的值再做一次。

// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
// 添加一个名为 sayHello 的新函数，它接受一个 Speaker 类型的值。
// 实现该函数以调用接口值上的 speak 方法。
// 然后创建每种类型的新值并使用该函数。

import "fmt"

type speaker interface {
	speak()
}

type english struct{}

func (english) speak() {
	fmt.Println("Hello")
}

type chinese struct{}

func (*chinese) speak() {
	fmt.Println("你好")
}

func main() {
	var s speaker
	s = english{}
	s.speak()

	s = &chinese{}
	s.speak()

	sayHello(english{})
	sayHello(&english{})
	sayHello(&chinese{})
	//sayHello(chinese{})		// chinese的值类型没有没有实现接口speaker
	// ./exercise.go:31:18: cannot use chinese{} (type chinese) as type speaker in argument to sayHello:
	//	chinese does not implement speaker (speak method has pointer receiver)
}

func sayHello(s speaker) {
	s.speak()
}

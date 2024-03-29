package main

import "fmt"

// Sample program to show how to declare and initialize struct types.
// 声明和初始化结构体类型

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// 初始化并赋予零值
	// Declare a variable of type example set to its zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

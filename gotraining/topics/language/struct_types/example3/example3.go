package main

import "fmt"

// Sample program to show how to declare and initialize anonymous
// struct types.
// 演示声明初始化匿名结构体类型变量

func main() {
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}
	// Display the value.
	fmt.Printf("%+v\n", e1)

	// 匿名结构体
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 1,
		pi:      3.1415,
	}

	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

package main

import "fmt"

// Sample program to show how variables of an unnamed type can
// be assigned to variables of a named type, when they are
// identical.

// 匿名类型的变量可赋值给相同显示类型的变量

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	var ex example
	// 未命名的类型赋给相同字段的显示类型.
	ex = e

	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", ex)
	fmt.Println("Flag", ex.flag)
	fmt.Println("Counter", ex.counter)
	fmt.Println("PI", ex.pi)
}

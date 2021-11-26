package main

import "fmt"

// Sample program to show the syntax and mechanics of type
// switches and the empty interface.
// 显示类型选择和空接口的语法和机制的示例程序。

// - 空接口没有说明接口内存储的数据。
// - 需要在运行时执行检查以了解有关存储在空接口中的数据的任何信息。
// - 围绕明确定义的行为解耦，仅在合理且可行的情况下将空接口用作例外。

func main() {
	// fmt.Println 可以传入任何参数
	fmt.Println("Hello")
	fmt.Println(123)
	fmt.Println(3.14159)
	fmt.Println(true)

	myPrintln("Hello")
	a := rune(10)
	myPrintln(a)
	myPrintln(123)
	myPrintln(3.14159)
	myPrintln(true)
}

func myPrintln(a interface{}) {
	// 类型断言语法
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string: type(%T): value(%s)\n", v, v)
	case int:
		fmt.Printf("Is string: type(%T): value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is string: type(%T): value(%f)\n", v, v)
	default:
		fmt.Printf("Is string: type(%T): value(%v)\n", v, v)
	}
}

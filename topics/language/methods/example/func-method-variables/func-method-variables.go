package main

import "fmt"

// Sample program to show how to declare function variables.
// 演示如何声明函数变量的示例程序。

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("Name is ", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "'age is ", d.age)
}

func main() {
	d := data{
		name: "Li",
	}
	fmt.Println("开始调用之前: ")
	d.displayName()
	d.setAge(45) // == (&d).setAge(45)

	fmt.Println("\n 编译过程:") // go 函数为值传递,参数为值则事复制副本为实参,参数为指针则为指针参数为实参
	data.displayName(d)
	(*data).setAge(&d, 45) // 编译器自动寻址.

	fmt.Println("\n 值方法变量调用:")
	// 值方法赋值给变量时,传递为副本而不是原值,所以期间变化的属性不是副本而是原值
	f1 := d.displayName
	f1()
	d.name = "Joan"
	f1()

	fmt.Println("\n 指针方法变量调用:")
	// 指针白方法赋值给变量时,传递的是原值的地址指针,所以期间变量的属性会一样会映射到指针中.
	f2 := d.setAge
	f2(45)
	d.name = "Sammy"
	f2(45)
}

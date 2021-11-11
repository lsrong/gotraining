package main

import "fmt"

// Sample program to show how the concrete value assigned to
// the interface is what is stored inside the interface.
// 示例程序显示分配给接口的具体值如何存储在接口内。
// 值实现: 存储值的副本,不会反应原始值的变化.
// 指针实现: 存储地址的服务,会反应原始值的变化.

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Printf("Printer Name: %s \n", c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Printf("Printer Name: %s \n", e.name)
}

func main() {
	// 定义两个结构体, canon为值接受者, epson为指针接受者, 都是实现接口printer
	c := canon{"PIXMA TR4520"}
	e := epson{"WorkForce Pro WR-3720"}

	// 存储到接口类型,
	printers := []printer{
		c,
		&e,
	}

	c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100"

	for _, p := range printers {
		p.print()
	}
	// Printer Name: PIXMA TR4520
	// Printer Name: Home XP-4100
	// canon: 存储一个值时，接口值有它自己的值副本。不会看到对原始值的更改。
	// *epson: 存储一个指针时，接口值有它自己的地址副本。将看到对原始值的更改。
}

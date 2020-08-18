package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// 变量声明的三种方式
	var a int
	var b = 10
	c := "string"
	a = 1000

	fmt.Printf("a=%d,b=%d,c=%s", a, b, c)

	fmt.Print("\r\n")

	// 多变量声明
	//只声明
	var d, e string
	d = "first"
	e = "second"
	fmt.Printf("d=%s,e=%s", d, e)

	fmt.Print("\r\n")

	// 声明并赋值
	f, g := 1, 2
	fmt.Printf("f=%d,g=%d", f, g)

	fmt.Print("\r\n")

	var (
		h string = "h string"
		i int    = 520
	)
	fmt.Printf("h=%s,i=%d", h, i)

	fmt.Print("\r\n")

	// 变量值互换
	m, n := 1, 2
	fmt.Printf("m=%d,n=%d", m, n)
	n, m = m, n
	fmt.Printf("m=%d,n=%d", m, n)

	//  _丢弃变量
	_ = n
	fmt.Printf("n=%d", n)
	fmt.Print("\r\n")
	// 多数据分组书写
	const (
		pi     = 3.1415
		prefix = "Go_"
	)

	// var (
	// 	i      int
	// 	pi     float32
	// 	prefix string
	// )
	fmt.Printf("%f,%s", pi, prefix)
	fmt.Print("\r\n")
	// 关键字iota
	// 关键字iota声明初始值为0，每行递增1：
	const (
		constA = iota // 0
		constB = iota // 1
		constC = iota // 2
	)
	fmt.Printf("%d%d%d", constA, constB, constC)
	fmt.Print("\r\n")
	const (
		constD = iota //  0
		constE        // 1
		constF        // 2
	)
	fmt.Printf("%d%d%d", constD, constE, constF)
	fmt.Print("\r\n")
	//如果iota在同一行，则值都一样
	const (
		constG                 = iota             //0
		constH, constI, constJ = iota, iota, iota // 1,1,1
		// k = 3                    // 此处不能定义缺省常量，会编译错误
	)
	fmt.Printf("%d%d%d%d", constG, constH, constI, constJ)

	fmt.Print("\r\n")
	// 结构体类型变量
	user := User{
		"LSRONG",
		28,
	}
	fmt.Printf("输出结构体变量：%+v", user)
}

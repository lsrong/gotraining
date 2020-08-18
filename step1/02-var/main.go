package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 说明
func information() {
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

// 变量
func testVar() {
	/**
	var a int
	var b bool
	var c string
	var d float32
	*/
	var (
		a int
		b bool
		c string
		d float32
	)

	fmt.Printf("a=%d,b=%t,c=%s,d=%f", a, b, c, d)

	a = 10
	b = true
	c = "hello"
	d = 0.10
	fmt.Printf("a=%d,b=%t,c=%s,d=%f", a, b, c, d)
}

// 常量
func testConst() {
	// 定义既赋值 const name [type] = value
	//const a int = 100
	//const b string = "hello"

	//fmt.Printf("a=%d", a)
	//fmt.Printf("b=%s", b)

	//const (
	//	c int    = 100
	//	d string = "world"
	//)

	//fmt.Printf("c=%d", c)
	//fmt.Printf("d=%s", d)

	const (
		a int = 100
		b
		c string = "hello"
		d
	)

	fmt.Printf("a=%d,b=%d,c=%s,d=%s", a, b, c, d)

	// Iota,默认为 0\
	/**
	const (
		e = iota
		f = iota
		g = iota
	)*/
	const (
		e = iota
		f
		g
	)
	fmt.Printf("e=%d,f=%d,g=%d", e, f, g)

	// <<< iota
	const (
		h = 1 << iota
		i
		j
	)
	fmt.Printf("h=%d,i=%d,j=%d", h, i, j)
}

func main() {
	//information()

	//testVar()

	//testConst()

	// iota 遇到非iota则会清零
	const (
		A = iota
		B
		C
		D = 8
		E
		F = iota
		G
	)
	// iota只在一个const生效
	const (
		A1 = iota
		A2
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(A1)
	fmt.Println(A2)
}

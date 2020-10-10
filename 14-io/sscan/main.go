package main

import "fmt"

// 从字符串中获取输入，格式化输入

//fmt.Sscanf(str, format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和
//格式化输出⼀致
func testSscanf() {
	var (
		a   int
		b   string
		c   float32
		str = "1 hello 1.0"
	)
	_, _ = fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a=%d,b=%s,c=%f", a, b, c)
}

//fmt.Sscan(str string, a …interface{}): 从终端获取⽤户输⼊，存储在Scanln中的参数⾥，
//空格和换⾏符作为分隔符
func testSscan() {
	var (
		a   int
		b   string
		c   float32
		str = "88 hello 8.8"
	)
	_, _ = fmt.Sscan(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

// fmt.Sscanln(str string, a …interface{}): 从终端获取⽤户输⼊，存储在Scanln中的参数⾥，
//空格作为分隔符，遇到换⾏符结束
func testSscanln() {
	var (
		a   int
		b   string
		c   float32
		str = "88 hello 8.8"
	)
	_, _ = fmt.Sscanln(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	//testSscanf()

	//testSscan()

	testSscanln()
}

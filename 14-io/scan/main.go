package main

import "fmt"

// fmt.Scan(a …interface{}): 从终端获取⽤户输⼊，存储在Scanln中的参数⾥，空格和换⾏符作为分隔符,回车结束
func testScan() {
	var (
		a int
		b string
		c float32
	)
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d,b=%s,c=%f", a, b, c)
}

// fmt.Scanf(format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和格式化输出⼀致
func testScanf() {
	var (
		a int
		b string
		c float32
	)
	_, _ = fmt.Scanf("%d\n", &a)
	_, _ = fmt.Scanf("%s\n", &b)
	_, _ = fmt.Scanf("%f\n", &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

// fmt.Scanln(a …interface{}): 从终端获取⽤户输⼊，存储在Scanln中的参数⾥，空格作为分隔符，遇到换⾏符结束
func testScanln() {
	var (
		a int
		b string
		c float32
	)
	_, _ = fmt.Scanln(&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	//testScan()

	//testScanf()

	testScanln()
}

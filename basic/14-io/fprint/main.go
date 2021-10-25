package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a int
		b string
		c float32
	)
	// 文件格式化输入
	_, _ = fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	//fmt.Printf("a=%d,b=%s,c=%f", a,b,c)

	// 写入文件
	_, _ = fmt.Fprintln(os.Stdout, "Stout", a, b, c)
}

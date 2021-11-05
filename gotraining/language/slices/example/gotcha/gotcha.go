package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

// 演示一个可能的陷阱示例]
// form: https://go.dev/blog/slices-intro
/*
重新切片切片不会复制底层数组。完整数组将保存在内存中，直到不再被引用。
有时这会导致程序在只需要一小部分数据时将所有数据保存在内存中。

*/
var digitRegexp = regexp.MustCompile("[0-9]+")

// FindDigit 返回文件里面[]byte指向包含整个文件的数组切片引用, 垃圾回收器不会销毁整个文件数组内存.
func FindDigit(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// 返回文件字节数组邻面
	return digitRegexp.Find(b)
}

// CopyDigit 将结果拷贝到一个新切片返回,这样大数据切片没有引用之后就可以被垃圾回收器及时销毁
func CopyDigit(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)

	return c
}

func main() {
	// 没有及时回收文件内容大数组内存
	filename := "demo.txt"
	digit1 := FindDigit(filename)
	fmt.Printf("digit1: %s\n", digit1)

	digit2 := CopyDigit(filename)
	fmt.Printf("digit2: %s\n", digit2)
}

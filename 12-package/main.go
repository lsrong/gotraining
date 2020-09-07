package main

import (
	"fmt"
	"github.com/learning_golang/12-package/calculator"
)

func init() {
	fmt.Println("This is main init")
}

/**
⾸字⺟⼤写表示可导出，⼩写表示私有。不能被外部的包访问
init 遵循先进后执行的规则，最底层的init先执行
*/
func main() {
	a := 100
	b := 200
	sum := calculator.Add(a, b)
	fmt.Printf("package calculator sum id %d \n", sum)
	sub := calculator.Sub(a, b)
	fmt.Printf("package calculator sub id %d \n", sub)
}

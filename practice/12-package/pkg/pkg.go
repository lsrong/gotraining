package main

import (
	"fmt"

	"github.com/learning_golang/practice/12-package/pkg/cal"
)

/*
*
总出init函数的执行顺序为先进后出的，从最下层的子包开始一层一层往上执行init里面的代码
*/
func init() {
	fmt.Println("this is main init function!")
}

func main() {
	// 分别调用 Add和Sub包
	a, b := 100, 50

	// Add
	c := cal.Add(a, b)
	fmt.Printf("call to [Add] function, ans= %d \n", c)

	// Sub
	d := cal.Sub(a, b)
	fmt.Printf("call to [Sub] function, ans= %d \n", d)
}

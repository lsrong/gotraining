package main

import (
	"fmt"
	"math"
)

/**
方法只是个带接收者参数的函数。
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// *** 上面定义个Abs的方法实际和下面定义的函数效果是一样的
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("方法调用v.Abs()的结果：%f \n", v.Abs())
	fmt.Printf("函数调用Abs(v)的结果：%f \n", Abs(v))
}

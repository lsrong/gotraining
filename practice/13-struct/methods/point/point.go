package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 上面的指针方法等同于下面的指针函数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	f := 0.11
	v1 := &Vertex{3, 4}
	v1.Scale(f)
	fmt.Printf("调用指针方法： %f \n", v1)

	v2 := &Vertex{3, 4}
	ScaleFunc(v2, f)
	fmt.Printf("调用指针函数： %f \n", v2)

}

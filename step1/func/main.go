package main

import "fmt"

func maths(a int, b int) (int, int, int, float32) {
	four := float32(a / b)
	fmt.Println(a / b)
	return a + b, a - b, a * b, four
}
func main() {

	var (
		a int = 520
		b int = 1314
	)

	sum, sub, mul, div := maths(a, b)

	fmt.Printf("多参数返回,加:%d,减%d,乘%d,除%f", sum, sub, mul, div)

}

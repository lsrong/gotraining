package main

// Sample program to show how to declare constants and their
// implementation in Go.
// 常量的基本定义和使用

// Constants live within the compiler.
// They have a parallel type system.
// Compiler can perform implicit conversions of untyped constants.
// 常量存在于编译器中。他们有一个并行类型系统。编译器可以执行无类型常量的隐式转换。

func main() {
	// 无类型常量
	const ui = 1234    // int
	const uf = 3.14159 // float64

	// 也可使用内置类型,有精度限制
	const ti int = 1234        // int
	const tf float64 = 5.14159 // float64

	// 不可超过类型取值范围
	//const muUint8 uint8 = 1000 // ./declare-init.go:22:8: constant 1000 overflows uint8

	// 常量算术运算支持不同类型, 隐式转换

	// 变量转换
	var answer = 3 * 3.33 // float64(3) * float64(3.33) = float64
	_ = answer

	// 常量也支持
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// kind will be integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3.0)

	// 无类型与有类型常量来执行数学运算,无类型会自动转换成有类型,结果为有类型常量的类型
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
}

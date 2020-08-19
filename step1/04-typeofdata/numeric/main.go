package main

import (
	"fmt"
	"math"
)

func testNumber() {
	/**
	1.整型:
	有符号:
	int     32位系统占4字节（与int32范围一样），64位系统占8个节（与int64范围一样）
	int8    占据1字节   范围 -128 ~ 127
	int16   占据2字节   范围 -2(15次方) ~ 2（15次方）-1
	int32   占据4字节   范围 -2(31次方) ~ 2（31次方）-1
	int64   占据8字节   范围 -2(63次方) ~ 2（63次方）-1
	rune	int32的别称

	无符号
	uint	32位系统占4字节（与uint32范围一样），64位系统占8字节（与uint64范围一样）
	uint8   占据1字节   范围 0 ~ 255
	uint16  占据2字节   范围 0 ~ 2（16次方）-1
	uint32  占据4字节   范围 0 ~ 2（32次方）-1
	uint64  占据8字节   范围 0 ~ 2（64次方）-1
	byte	uint8的别称
	*/

	/**
	2.浮点类型:
	float32 单精度  占据4字节   范围 -3.403E38 ~ 3.403E38    (math.MaxFloat32)
	float64 双精度  占据8字节   范围 -1.798E208 ~ 1.798E308  (math.MaxFloat64)
	*/
	fmt.Println(math.MaxFloat32)

	fmt.Println(math.MaxFloat64)

	/**
	注意:精度缺失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1=",num1)		// -123.00009
	fmt.Println("num2=",num2)		// -123.0000901
	*/
	var num1 float64 = -123.0000901
	var num2 float64 = -123.000901
	fmt.Println("num1=", num1) // -123.00009
	fmt.Println("num2=", num2) // -123.0000901
	fmt.Println(isFloatEqual(num1, num2, 0.0001))

	/**
	3.NaN
	*/
	nan := math.NaN()
	fmt.Println(math.IsNaN(nan))
}

// 浮点型比较
func isFloatEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}

func testInt() {
	var a int8
	a = 100
	fmt.Println(a)

	var b int32 = 100

	// 不同数据类型不允许相加，以及赋值操作
	// 如果需要则需要进行数据转换， typeOfData()

	fmt.Println(int32(a) + b)

	// 占位符： %d , %f
}

func main() {
	testNumber()

	testInt()
}

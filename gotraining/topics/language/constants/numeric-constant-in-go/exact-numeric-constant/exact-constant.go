package main

import "fmt"

// 比 int64 大得多的值。
//const myConst = 9223372036854775808543522345

// constant 9223372036854775808543522345 overflows int64
const myConst int64 = 9223372036854775808543522345

func main() {
	//fmt.Println("将编译")

	fmt.Println("Will NOT Compile")
}

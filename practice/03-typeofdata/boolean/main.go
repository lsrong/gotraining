package main

import (
	"fmt"
)

func main() {
	var a bool

	fmt.Println(a)

	a = true
	fmt.Println(a)

	// !
	a = !a
	fmt.Println(a)

	// &&
	a = true
	var b bool
	if a == true && b == true {
		fmt.Println("It is true")
	} else {
		fmt.Println("It is false")
	}

	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}

	// %t 格式化输出
	fmt.Printf("%t %t \n", a, b)

}

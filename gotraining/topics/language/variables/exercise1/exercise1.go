// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
// 声明三个初始化为零值的变量和三个用文字值声明的变量。
// 声明类型为 string、int 和 bool 的变量。
// 显示这些变量的值。
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
// 声明一个 float32 类型的新变量，并通过转换 Pi (3.14) 的文字值来初始化该变量。
package main

import "fmt"

func main() {
	// Declare variables
	var age int
	var name string
	var legal bool
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	fmt.Println("---")

	// Declare variables and initialize.
	month := 10
	dayOfWeek := "Wednesday"
	happy := true
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	fmt.Println("---")

	// Conversion
	pi := float32(3.14159)
	fmt.Printf("%T [%v] \n", pi, pi)
}

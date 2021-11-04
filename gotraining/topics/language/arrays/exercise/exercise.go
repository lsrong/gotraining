package main

import "fmt"

// Declare an array of 5 strings with each element initialized to its zero value.
// 声明一个包含 5 个字符串的数组，每个元素都初始化为零值。
//
// Declare a second array of 5 strings and initialize this array with literal string values.
// 声明第二个由 5 个字符串组成的数组，并用文字字符串初始化这个数组值。
// Assign the second array to the first and display the results of the first array.
// 将第二个数组分配给第一个数组并显示第一个数组的结果。
// Display the string value and address of each element.
// 显示每个元素的字符串值和地址。

func main() {
	var names [5]string

	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	names = friends
	for i, name := range names {
		fmt.Printf("Value[%s]\t Address[%p] \n", name, &names[i])
	}
}

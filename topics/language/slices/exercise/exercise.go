package main

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
// 声明一个 nil 整数切片。创建一个循环，将 10 个值附加到切片。迭代切片并显示每个值。
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slic
// 声明一个包含五个字符串的切片，并使用字符串文字值初始化该切片。
// 显示所有元素。取索引一、二的切片，显示新切片中每个元素的索引位置和值.

import "fmt"

func main() {
	var digits []int

	for i := 0; i < 10; i++ {
		digits = append(digits, i)
	}

	fmt.Printf("%+v\n", digits)
	for _, number := range digits {
		fmt.Println(number)
	}

	names := []string{"Bill", "Ming", "Li", "Jenny"}
	for i, name := range names {
		fmt.Printf("Index[%d]\tName[%s]\n", i, name)
	}

	peoples := names[1:3]
	for i, people := range peoples {
		fmt.Printf("Index[%d]\tPeople[%s]\n", i, people)
	}

}

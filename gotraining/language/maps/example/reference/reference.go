package main

import "fmt"

// Sample program to show how maps are reference types.

// * Maps are a reference type.
//* map是一种引用类型。

func main() {
	scores := map[string]int{
		"Li": 45,
		"Ye": 26,
	}

	// map参数传递的 指针值，
	double(scores, "Ye")

	fmt.Println("Score: ", scores["Ye"])
}

// double 映射参数为指针参数。
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}

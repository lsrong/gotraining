package main

import "fmt"

func main() {

	colors := []string{"Red", "Blue", "White", "Black"}
	// 值遍历,如果原切片变化不会影响值遍历
	for _, v := range colors {
		colors = colors[:2]
		fmt.Printf("v[%s]\n", v)
	}

	colors = []string{"Red", "Blue", "White", "Black"}
	// 指针遍历,如果原切片变化会影响到遍历结果(直接读原数据)
	for i, _ := range colors {
		//colors = colors[:2]
		fmt.Printf("v[%s]\n", colors[i])

		// panic: runtime error: index out of range [2] with length 2
	}

	demoFilterFunc := func(s string) bool {
		return s == "Blue" || s == "Black"
	}

	colors = Filter(colors, demoFilterFunc)

	fmt.Println(colors)
}

func Filter(slice []string, f func(s string) bool) []string {
	r := slice[:0]
	for _, v := range slice {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

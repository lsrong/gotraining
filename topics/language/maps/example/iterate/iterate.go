package main

import "fmt"

// Sample program to show how to declare, initialize and iterate
// over a map. Shows how iterating over a map is random.

//* Iterating over a map is always random.
//* 迭代map总是随机的。

// maps 的遍历是无序的

type user struct {
	id   int
	name string
}

func main() {
	// 定义并初始化map
	users := map[int]user{
		30: {id: 30, name: "Sun"},
		10: {id: 10, name: "LI"},
		20: {id: 20, name: "Zhang"},
	}

	// （无序）遍历
	for k, u := range users {
		fmt.Println(k, u)
	}

	// range只有一个键遍历的是key
	for k := range users {
		fmt.Println(k)
	}
}

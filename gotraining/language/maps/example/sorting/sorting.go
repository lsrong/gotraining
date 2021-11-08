package main

import (
	"fmt"
	"sort"
)

// Sample program to show how to walk through a map by
// alphabetical key order.

// 通过获取key在排序来实现有序遍历的

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
		8:  {id: 8, name: "YoYo"},
	}

	// 获取所有key
	var keys []int
	for k := range users {
		keys = append(keys, k)
	}

	// 使用sort扩展库排序
	sort.Ints(keys)

	// 遍历排序后的键切片来打印map的值
	for _, k := range keys {
		fmt.Println(k, users[k])
	}

}

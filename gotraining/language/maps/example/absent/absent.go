package main

import "fmt"

// Sample program to show how maps behave when you read an
// absent key.

// 不存在的键值处理方式

//* Reading an absent key returns the zero value for the map's value type.
//* 读取不存在的键会返回映射值类型的零值。

func main() {
	// 简单模拟分数计数
	scores := make(map[string]int)

	// 不存在的值返回类型的零值
	score := scores["anna"]
	fmt.Println("Score: ", score)

	// 读取的第二个参数为bool类型, 是否存在
	score, ok := scores["anna"]
	fmt.Println("Score: ", score, "Present: ", ok)

	// 可以利用零值来方便直接操作不存在的哈希值
	scores["anna"]++

	// 如果没有零值机制就需要用以下机制来防止没有初始化情况.
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score: ", score, "Present: ", ok)
}

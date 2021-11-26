package main

import "fmt"

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.

// 声明并制作一个以字符串为键的整数值映射。用五个值填充地图并迭代地图以显示键值对。

func main() {
	players := make(map[string]int)

	players["Li"] = 45
	players["Ye"] = 36
	players["Ming"] = 33

	for player, score := range players {
		fmt.Printf("Player: %s \t Score: %d \n", player, score)
	}
}

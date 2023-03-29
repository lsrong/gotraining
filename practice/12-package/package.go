package main

import (
	"fmt"

	"github.com/learning_golang/practice/12-package/calculator"
)

func init() {
	fmt.Println("This is main init")
}

// --------- 练习题
/**
你有50枚⾦币，需要分配给以下⼏个⼈：Matthew, Sarah, Augustus, Heidi, Emilie,
 Peter, Giana, Adriano, Aaron, Elizabeth。分配规则如下所示：
a. 名字中包含’a’或’A’: 1枚⾦币
b. 名字中包含’e’或’E’: 1枚⾦币
c. 名字中包含 ‘i’或’I’: 2枚⾦币
d. 名字中包含’o’或’O’: 3枚⾦币
e. 名字中包含’u’或’U’: 5枚⾦币

 写⼀个程序，计算每个⽤户分到了多少⾦币，以及最后剩余多少⾦币？
*/
var (
	users = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distributions = make(map[string]int, len(users))
)

/*
*
首字符大写表示可导出，小写表示私有化，不能被外部的包访问
*/
func main() {
	// init 遵循先进后出的规则，最底层的init先执行。
	a, b := 10, 20
	_ = calculator.Add(a, b)
	_ = calculator.Sub(a, b)

	// 演示分配金币
	coins := 50
	fmt.Printf("Before allocate coins=%d \n", coins)
	allocate(&coins)
	fmt.Printf("After allocate coins=%d \n", coins)
	for username, coin := range distributions {
		fmt.Printf("The user[%s] receives %d coins\n", username, coin)
	}
}

// receiveCoin 每个字符可以分配的金币数（switch）
func receiveCoin(char byte) int {
	coin := 0
	switch char {
	case 'a', 'A', 'e', 'E':
		coin = 1
	case 'i', 'I':
		coin = 2
	case 'o', 'O':
		coin = 3
	case 'u', 'U':
		coin = 5
	}
	return coin
}

// allocate 遍历每个名字的字符出现次数，并统计金币数量
func allocate(coins *int) {
	for _, username := range users {
		for _, char := range username {
			coin := receiveCoin(byte(char))
			v, ok := distributions[username]
			if ok {
				distributions[username] = v + coin
			} else {
				distributions[username] = v
			}
			*coins -= coin
		}
	}
}

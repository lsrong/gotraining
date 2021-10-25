package main

import (
	"fmt"
)

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

func receiveCoin(char byte) int {
	coin := 0
	switch char {
	case 'a', 'A', 'e', 'E':
		coin = 1
	case 'i', 'I':
		coin = 2
	case 'O', 'o':
		coin = 3
	case 'u', 'U':
		coin = 5
	}
	return coin
}

func allocate(coins *int) {
	for _, username := range users {
		for _, char := range username {
			coin := receiveCoin(byte(char))
			value, ok := distributions[username]
			if ok {
				distributions[username] = value + coin
			} else {
				distributions[username] = coin
			}
			*coins -= coin
		}
	}
}

func main() {
	coins := 100
	fmt.Printf("Before allocate coins=%d \n", coins)
	allocate(&coins)
	fmt.Printf("After allocate coins=%d \n", coins)
	for username, coin := range distributions {
		fmt.Printf("The user[%s] receives %d coins \n", username, coin)
	}
}

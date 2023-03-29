package main

import "fmt"

/*
*
你有50枚⾦币，需要分配给以下⼏个⼈：Matthew, Sarah, Augustus, Heidi, Emilie,
Peter, Giana, Adriano, Aaron, Elizabeth。分配规则如下所示：
a. 名字中包含’a’或’A’: 1枚⾦币
b. 名字中包含’e’或’E’: 1枚⾦币
c. 名字中包含 ‘i’或’I’: 2枚⾦币
d. 名字中包含’o’或’O’: 3枚⾦币
e. 名字中包含’u’或’U’: 5枚⾦币

	写⼀个程序，计算每个⽤户分到了多少⾦币，以及最后剩余多少⾦币？
*/
var peoples = []string{
	"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
}

func main() {
	glod := 100
	ps := assign(&glod)
	for p, g := range ps {
		fmt.Printf("用户【%s】分到了 %d 金币\n", p, g)
	}
	fmt.Printf("最后剩余 %d 个金币！", glod)

}

func assign(glod *int) map[string]int {
	ps := make(map[string]int, len(peoples))
	for _, p := range peoples {
		g := allocate(p)
		v, ok := ps[p]
		if !ok {
			ps[p] = g
		} else {
			ps[p] = v + g
		}
		*glod -= g
	}
	return ps
}

func allocate(name string) int {
	glod := 0
	for _, v := range []byte(name) {
		switch v {
		case 'a', 'A', 'e', 'E':
			glod += 1
		case 'i', 'I':
			glod += 2
		case 'o', 'O':
			glod += 3
		case 'u', 'U':
			glod += 5
		}
	}
	return glod
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	demoMap()
	fmt.Println("华丽的分割线--------")
	demoSortMap(10)
}

func demoMap() {
	// 赋值并初始化
	a := map[string]string{
		"hello": "你好",
		"world": "世界",
	}
	fmt.Printf("hello world! \n%s %s \n", a["hello"], a["world"])

	// map类型的默认初始化为nil, 要使用make分配map内存
	var b map[string]int
	// b["go"] = 100 // assiggnment to entry in nil map
	b = make(map[string]int)
	b["go"] = 100
	fmt.Printf("b[%s] = %d \n", "go", b["go"])

	// 查：判断key是否存在
	var c = map[string]int{
		"go":   100,
		"php":  80,
		"java": 60,
	}
	_, ok := c["go"]
	fmt.Printf("go是否存在[%v] \n", ok)
	// 不存在才添加
	_, ok = c["C++"]
	if !ok {
		c["C++"] = 40
	}

	// 遍历: for k, v :=range map
	for k, v := range c {
		fmt.Printf("语言【%s】weight=%d \n", k, v)
	}

	// 删除 delete(map)
	delete(c, "C++")
	fmt.Printf("C++不适合我，已经删除\n")
}

// demoSortMap 默认情况下，map并不是按照key有序进行遍历的
// 换一种方式排序，键值数组进行排序，根据键来输出mao值
func demoSortMap(num int) {
	// 构建map变量
	demoMap := make(map[int]string, num)
	for i := 1; i <= num; i++ {
		demoMap[i] = fmt.Sprintf("v_%d", i)
	}

	// 去除所有的键进行排序
	keys := []int{}
	for i, _ := range demoMap {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	// 按照顺序打印
	for _, key := range keys {
		fmt.Printf("demoMap[%d] = %s \n", key, demoMap[key])
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	demoMap()
	demoSortMap()
}

// demoMap map的增删改查 key-value, map必须初始化之后才能使用
func demoMap() {
	// 初始化并赋值
	var a = map[string]string{
		"hello": "world",
		"hi":    "hello",
	}
	fmt.Printf("%s %s \n", a["hi"], a["hello"])

	// map类型的默认初始化为nil, 要使用make分配map内存
	var b map[string]int
	// b["key"] = 10 //报错:assignment to entry in nil map
	b = make(map[string]int, 2) // 要使用make分配内存
	b["a"] = 1
	b["b"] = 2
	fmt.Printf("make demo b=%v\n", b)

	// 查：判断key是否存在
	c := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	v, ok := c["a"]
	fmt.Printf("whether is exist a:%v\n", ok)
	fmt.Printf("the value of a is %d\n", v)
	v, ok = c["d"]
	if !ok {
		v = 4
		c["d"] = v
	}

	// 遍历: for k, v :=range map
	var langs = map[string]int{
		"go":   100,
		"java": 80,
		"php":  60,
	}
	for k, v := range langs {
		fmt.Printf("a[%s] = %d \n", k, v)
	}

	// 删除 delete(map)
	delete(langs, "php")
	for k, v := range langs {
		fmt.Printf("a[%s] = %d \n", k, v)
	}
	fmt.Printf("map langs length is %d \n", len(langs))
}

// demoSortMap 默认情况下，map并不是按照key有序进行遍历的
// 换一种方式排序，键值数组进行排序，根据键来输出mao值
func demoSortMap() {
	// 随机生成，map
	rand.Seed(time.Now().UnixNano())
	l := 10
	a := make(map[int]string, l)
	for i := 0; i < l; i++ {
		a[i] = fmt.Sprintf("v_%d", i)
	}

	keys := make([]int, 0, l)
	for i := range a {
		fmt.Printf("map[%d]=%s \n", i, a[i])
		keys = append(keys, i)
	}
	fmt.Println("------")
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("map[%d]=%s \n", key, a[key])
	}
}

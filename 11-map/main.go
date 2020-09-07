package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// key-value的数据结构, map必须初始化才能使⽤
func initMap() {
	var a = map[string]string{
		"name":  "hello",
		"value": "word",
	}

	fmt.Println(a)

	key := "value"
	fmt.Printf("the value of  key[name] is %s \n", a["name"])
	fmt.Printf("the value of  key[%s] is %s \n", key, a[key])
}

// map类型的变量默认初始化为nil，需要使⽤make分配map内存
func nilMap() {
	var a map[string]int
	//a["key01"] = 100
	//a["key02"] = 200
	fmt.Println(a)
	a = make(map[string]int, 2)
	a["key01"] = 100
	a["key02"] = 200
	fmt.Println(a)
}

// 判断map指定的key是否存在: value, ok := map[key]
func okMap() {
	a := map[string]int{
		"key01": 1,
		"key02": 2,
		"key03": 3,
	}
	fmt.Println(a)

	v1, ok := a["key01"]
	fmt.Printf("whether is exist key01:%v\n", ok)
	fmt.Printf("the value of key01 is %d\n", v1)

	v4, ok := a["key04"]
	if ok {
		fmt.Println("the key[key04] is exist")
	} else {
		fmt.Println("the key[key04] is not exist")
		v4 = 4
		a["key04"] = v4
	}

	fmt.Println(a)
}

// map遍历操作:for k,v := range map {}
func forMap() {
	var language = map[string]int{
		"go":     100,
		"python": 90,
		"java":   80,
		"c++":    70,
		"php":    60,
	}
	for k, v := range language {
		fmt.Printf("a[%s] = %d\n", k, v)
	}
}

// map删除元素
func deleteMap() {
	a := make(map[int]string)
	a[0] = "hello"
	a[1] = " word"
	a[2] = "!!"
	fmt.Println(a)
	delete(a, 2)
	fmt.Println(a)
	for k := range a {
		delete(a, k)
	}
	fmt.Println(a)
}

// map的⻓度
func lenMap() {
	a := map[int]string{
		0: "hello",
		1: "world",
		2: "go",
	}
	a[3] = "test"
	fmt.Println("map length is ", len(a))
}

// map是引⽤类型
func modifyMap(a map[string]string) {
	a["hello"] = "world"
}
func testModifyMap() {
	a := map[string]string{
		"test": "content",
	}
	fmt.Println("Before modify ", a)
	modifyMap(a)
	fmt.Println("After modify ", a)
}

// 默认情况下，map并不是按照key有序进⾏遍历
// 建立临时keys数组，键值数组排序，然后循环根据键数组输出map值
func sortMap() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[int]string, 10)
	for i := 0; i < 10; i++ {
		a[i] = fmt.Sprintf("value_%d", i)
	}
	//fmt.Println(a)
	keys := make([]int, 0, 10)
	for i, v := range a {
		fmt.Printf("map[%d]=%s\n", i, v)
		keys = append(keys, i)
	}

	sort.Ints(keys)
	for _, mapKey := range keys {
		fmt.Printf("map[%d]=%s\n", mapKey, a[mapKey])
	}
}

func main() {
	initMap()
	fmt.Println("------------")
	nilMap()
	fmt.Println("------------")
	okMap()
	fmt.Println("------------")
	forMap()
	fmt.Println("------------")
	deleteMap()
	fmt.Println("------------")
	lenMap()
	fmt.Println("------------")
	testModifyMap()
	fmt.Println("------------")
	sortMap()
}

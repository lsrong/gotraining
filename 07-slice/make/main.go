package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建切片， make(type, len(长度), cap（容量）)
func testMake() {
	a := make([]int, 3, 5)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	// append
	a = append(a, 4)
	fmt.Println(a)

	// 自动扩容
	for i := 0; i < 10; i++ {
		a = append(a, i+10)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	fmt.Println(a)

	//cap 切片的容量
	b := a[:cap(a)]
	fmt.Println(b)
}

// 追加多个元素
func testAppend() {
	a := []int{0, 1, 2, 3, 4}
	var b []int
	b = append(b, a...)

	fmt.Println(b)
}

// 切片相加
func sumSplice(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func main() {
	testMake()

	testAppend()

	// 随机切片 => 切片相加
	rand.Seed(time.Now().Unix())
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)
	fmt.Println(sumSplice(a[:]))
}

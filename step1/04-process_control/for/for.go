package main

import (
	"fmt"
	"time"
)

// for基本用法
// for initialisation; condition;post{}
func testBaseFor() {
	for i := 1; i <= 10; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

		time.Sleep(time.Second)
	}
}

// 跳出循环，break
func testBreak() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			fmt.Println("The num is plus 5")
			break
		}
		fmt.Printf("i=%d \n", i)
	}
}

// 中止执行当前循环Continue
func testContinue() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Printf("i=%d is even \n", i)
	}
}

//  initialisation  for condition { post }  (类似于while)
func testWhile() {
	i := 0       //Initialisation
	for i < 10 { // Condition
		fmt.Printf("i=%d \n", i)
		i++ // Post
	}
}

// 同时赋值for循环
func testMultiple() {
	//var a int
	//var b string
	//a = 10
	//b = "hello"
	//fmt.Println(a, b)

	//a, b := 10, "hello"
	//fmt.Println(a, b)
	for i, j := 9, 1; i > 0 && j < 10; i, j = i-1, j+1 {
		fmt.Printf("%d*%d=%d \n", i, j, i*j)
	}
}

// 无限循环
func testUnlimited() {
	for {
		fmt.Println("hello")
	}
}

// 输出9*9乘法表
func multiplicationTable() {
	fmt.Println()
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	testBaseFor()

	testBreak()

	testContinue()

	testWhile()

	testMultiple()

	//testUnlimited()

	multiplicationTable()
}

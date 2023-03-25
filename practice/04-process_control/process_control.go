package main

import (
	"fmt"
	"time"
)

func main() {
	testIf()
	testSwitch()
	testFor()
}

// testIf  if - 判断语句
func testIf() {
	//基本 if condition {} [else if condition{} ...] else {}
	test1 := func(num int) {
		if num%2 == 0 {
			fmt.Printf("The number is even: %d \n", num)
		} else {
			fmt.Printf("The number is odd: %d \n", num)
		}
	}
	test1(99)

	// 赋值加判断  if assignment;condition {}[else if condition{}...] else {}
	test2 := func(num int) {
		if temp := num + 1; temp%2 == 0 {
			fmt.Printf("The num+1 is even \n")
		} else if temp := num - 1; temp%2 == 0 {
			fmt.Printf("The num-1 is even \n")
		} else if num%2 == 0 {
			fmt.Printf("The num is even \n")
		} else {
			fmt.Printf("The num is odd \n")
		}
	}
	test2(99)
	getN := func() int {
		return 99
	}
	// 执行函数 + 判断
	test3 := func() {
		if num := getN(); num%2 == 0 {
			fmt.Println("The value of 'get()' is even")
		} else {
			fmt.Println("The value of 'get()' is odd")
		}
	}
	test3()
}

func testSwitch() {
	// switch var {case condition: // 代码块 }
	a := 4
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	case 4:
		fmt.Println("a=4")
	case 5:
		fmt.Println("a=5")
	}
	value := func() int {
		return 10
	}
	//  switch assignment; var {case condition: // 代码块}
	testSwitchAssignment := func() {
		switch a := value(); a {
		case 1:
			fmt.Println("a=1")
		case 2:
			fmt.Println("a=2")
		case 3:
			fmt.Println("a=3")
		case 4:
			fmt.Println("a=4")
		case 5:
			fmt.Println("a=5")
		default:
			fmt.Println("a is plus 5")
		}
	}
	testSwitchAssignment()

	//  case 多个条件
	switch a := value(); a {
	case 1, 2, 3, 4, 5:
		fmt.Println("a >= 1 and a <= 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("a >= 6 and a <= 10")
	default:
		fmt.Println("a <=0 or a > 10")
	}

	// case 范围
	testRangeSwitch := func(num int) {
		switch {
		case num >= 0 && num <= 25:
			fmt.Println("num >= 0 and num <= 25")
		case num > 25 && num <= 50:
			fmt.Println("num > 25 and num <= 50")
		case num > 50 && num <= 75:
			fmt.Println("num > 50 and num <= 75")
		case num > 75 && num <= 100:
			fmt.Println("num > 75 and num <= 100")
		default:
			fmt.Println("The num is not in (0,100)")
		}
	}
	testRangeSwitch(80)

}

// testFor for - 循环语句
func testFor() {
	// for 基本用法  for init; cond; pos { // 代码块 }
	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}

	// break 跳出循环
	for i := 1; i < 10; i++ {
		if i > 5 {
			fmt.Println("The num is plus 5. Break!")
			break
		}
		fmt.Printf("i = %d \n", i)
	}
	// continue 跳出当前循环
	for i := 0; i < 10; i++ {
		if i%2 != 0 { // 取余
			continue
		}
		fmt.Printf("i=%d is even! \n", i)
	}

	// "while"  for cond  { pos }
	j := 0
	for j < 10 { // cond
		fmt.Printf("j=%d \n", j)
		j++ // post
	}

	// 同时赋值 - for
	for m, n := 9, 1; m > 0 && n < 10; m, n = m-1, n+1 {
		fmt.Printf("%d + %d = %d \n", n, m, n+m)
	}

	// 9*9乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 无限循环
	for {
		fmt.Println("HELLO")
		time.Sleep(time.Second)
	}
}

package main

import (
	"errors"
	"fmt"
)

var num int = 1

func echoLine() {
	fmt.Print("\r\n")
}

func main() {
	// 1.判断if
	param := 80
	if param >= 60 {
		fmt.Print("考试成绩大于60，及格")
	}

	echoLine()

	// 初始化与判断写一起
	if a := 10; a == 10 {
		fmt.Print("初始化与判断写在一起： if a := 10; a == 10")
	}

	echoLine()
	// 2.分支语句switch
	/*
		注意：
			Go保留了break，用来跳出switch语句，上述案例的分支中默认就书写了该关键字
			Go也提供fallthrough，代表不跳出switch，后面的语句无条件执行
	*/
	switch num {
	case 1:
		fmt.Print("111")
		// fallthrough
	case 2:
		fmt.Print("222")
	default:
		fmt.Print("000")
	}

	// 3.循环语句
	// 传统模式
	/*
		for int; condition; pos {

		}*/
	echoLine()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		echoLine()
	}
	// for循环简化
	var j int
	for ; ; j++ {
		if j >= 10 {
			break
		}
		fmt.Printf("%d", j)
		echoLine()
	}
	// 类似while循环
	var m int
	for m < 10 {
		fmt.Printf("%d", m)
		echoLine()
		m++
	}
	// 死循环
	/*
		for{

		}
	*/
	// for range:一般用于遍历数组、切片、字符串、map、管道
	for key, value := range []int{1, 2, 3} {
		fmt.Printf("遍历数组元素：key = %d,value=%d", key, value)
		echoLine()
	}

	// 3、跳出循环
	/*
		break:用于函数内跳出当前for、switch、select语句的执行,跳出当前循环体
		continue:用于跳出for循环的本次迭代。
		goto:可以退出多层循环
	*/
	// break:仅退出当前for,switch,select
	for i := 0; i < 2; i++ {
		for j = 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break
			case 3:
				fmt.Println(i, j)
				break
			}
		}
	}

	// goto:退出多层循环语句
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 6 {
				goto breakHere
			}
		}
	}
	// todo:不知道啥意思，待了解
breakHere:
	fmt.Println("goto break")

	// goto:统一错误处理
	var err = errors.New("hi")
	if err != nil {
		goto onExit
	}
onExit:
	fmt.Print("错误处理")
	fmt.Println(err)
}

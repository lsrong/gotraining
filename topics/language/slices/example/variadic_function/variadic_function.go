package main

// Sample program to show how to declare and use variadic functions.

// 演示声明和使用可变参数函数。
// 可变参数实际就是参数切片,如果实参也为切片的话就可以共享底层数组.

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	u1 := user{
		id:   1,
		name: "Betty",
	}
	u2 := user{
		id:   2,
		name: "Janet",
	}

	display(u1, u2)

	u3 := []user{
		{24, "Bill"},
		{42, "Joan"},
	}

	display(u3...)

	fmt.Println("**************************")
	change(u3...)
	for _, u := range u3 {
		fmt.Printf("%#+v\n", u)
	}

}

func display(users ...user) {
	for _, u := range users {
		fmt.Printf("%#+v\n", u)
	}
}

// 共享底层数组方式,可变参数(切片)
func change(users ...user) {
	users[1] = user{99, "change backing"}
}

package main

import "fmt"

// Sample program to show some of the mechanics behind the
// short variable declaration operator redeclares.
// 短变量声明运算符重新声明。

// From Spec:
// a short variable declaration may redeclare variables provided they
// were originally declared earlier in the same block with the same
// type, and at least one of the non-blank variables is new.

// 简短的变量声明可以重新声明变量，前提是它们最初是在同一块中以相同类型声明的，
// 并且至少有一个非空变量是新的。

type user struct {
	id   int
	name string
}

func main() {
	var err1 error
	// 定义返回新变量u, 重定义变量err1
	u, err1 := getUser(1, "bill")
	if err1 != nil {
		return
	}
	fmt.Println(u)

	// 重定义变量u, 新变量err2
	u, err2 := getUser(2, "Jenny")
	if err2 != nil {
		return
	}

	fmt.Println(u)
}

func getUser(id int, name string) (*user, error) {
	return &user{id: id, name: name}, nil
}

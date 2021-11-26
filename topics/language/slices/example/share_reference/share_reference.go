package main

import (
	"fmt"
)

// Sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.
// 示例程序显示当您引用元素时附加到切片时需要注意的事项。

// 追加的时候,如果超过容量时,底层的动作为:1,生成长度比较大的同类型新数组,2.复制原数组到新数组,3.将引用转移到新数组.
// 因此引用原数组的指针变量不会作用在对新扩容的元素上.

type user struct {
	likes int
}

func main() {
	users := make([]user, 3)
	shareUser := &users[1]
	shareUser.likes++
	for i := range users {
		fmt.Printf("User: %d likes: %d Address:%p\n", i, users[i].likes, &users[i])
	}

	fmt.Println("*************************")

	users = append(users, user{})
	shareUser.likes++
	for i := range users {
		fmt.Printf("User: %d likes: %d Address:%p\n", i, users[i].likes, &users[i])
	}
	// 最后的shareUser.likes不会在users切片中生效,因为底层数组已经变化
}

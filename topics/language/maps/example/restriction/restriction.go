package main

import "fmt"

// Sample program to show how only types that can have
// equality defined on them can be a map key.
// 具有可比较的类型才能作为键使用:boolean, numeric, string, pointer, channel, interface, structs or arrays
// 不可比较类型:slices, maps, and functions;

//* The map key must be a value that is comparable.
//* 映射键必须是可比较的值。

type user struct {
	id   int
	name string
}
type users []user

func main() {
	u := make(map[users]int) // ./restriction.go:20:12: invalid map key type users

	for k, v := range u {
		fmt.Println(k, v)
	}
}

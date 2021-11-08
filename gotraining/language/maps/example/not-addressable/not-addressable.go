package main

// Sample program to show that you cannot take the address
// of an element in a map.

//* Elements in a map are not addressable.
// map的值是不可寻址的。

type player struct {
	name  string
	score int
}

func main() {
	players := map[string]player{
		"li": {name: "li", score: 41},
		"yo": {name: "yo", score: 21},
	}

	// map的值是不可寻址的
	li := &players["li"] // ./not-addressable.go:20:8: cannot take the address of players["li"]
	li.score++

	// 操作方式： 取出元素，修改操作，放回, map
	pl := players["li"]
	pl.score++
	players["li"] = pl
}

package main

import "fmt"

// Sample program to show how to initialize a map, write to
// it, then read and delete from it

// 初始化map类型, 读,写,删除操作

type user struct {
	name    string
	surname string
}

func main() {

	// Declare and make a map variable.
	users := make(map[string]user)
	// or
	//users := map[string]user{}

	// add
	users["Li"] = user{name: "Li", surname: "Ming"}
	users["Zhang"] = user{name: "Zhang", surname: "Zhang"}
	users["Mouse"] = user{name: "Mickey", surname: "Mouse"}

	// Read
	mouse := users["Mouse"]
	fmt.Printf("%+v \n", mouse)

	// Replace
	users["Mouse"] = user{name: "Jerry", surname: "Mouse"}
	fmt.Printf("%+v \n", users["Mouse"])

	// delete
	delete(users, "Li")
	fmt.Println(len(users))

	// 删除不存在的键是安全的
	delete(users, "Li")
}

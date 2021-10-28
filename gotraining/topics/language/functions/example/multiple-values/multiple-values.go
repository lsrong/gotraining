package main

import (
	"encoding/json"
	"fmt"
)

// Sample program to show how functions can return multiple values while using
// named and struct types.
// 多返回值函数

var nextID = 1

type user struct {
	ID   int
	Name string
}

func main() {
	u, err := retrieveUser("bill")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", u)
}

// retrieveUser retrieves a pointer to user and error
func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}
	var u user
	err = json.Unmarshal([]byte(r), &u)

	return &u, err
}

func getUser(name string) (string, error) {
	response := fmt.Sprintf("{\"id\":%d, \"name\":\"%s\"}", nextID, name)

	return response, nil
}

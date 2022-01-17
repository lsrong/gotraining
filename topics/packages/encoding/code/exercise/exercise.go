// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// Marshal the slice into pretty print strings and display each element.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// User 用户信息
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 打开文件
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("open file: %v", err)
	}

	// 关闭文件
	defer file.Close()

	// 序列化userJSON
	var users []User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		log.Fatalf("Decode json: %v", err)
	}
	// 遍历用户s
	for i, user := range users {
		fmt.Printf("user[%d]: %+v \n", i, user)
	}

	// Marshal each user value and display the JSON. Check for errors.
	uData, err := json.MarshalIndent(&users, "", "    ")
	if err != nil {
		log.Fatalf("MarshalIndent: %v", err)
	}

	fmt.Println(string(uData))
}

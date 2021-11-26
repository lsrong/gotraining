package main

// Sample program to show how we can use the blank identifier to
// ignore return values.

// 使用空白标识符`_`来忽略返回值.

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	ID   int
	Name string
}

// updateStats 更新数据
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	u := &user{
		ID:   1,
		Name: "Jenny",
	}

	// 不关系返回接口可以忽略返回变量
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("updated user record for ID ", u.ID)
}

func updateUser(u *user) (*updateStats, error) {
	//response := `{"Modified":1, "Duration":0.005, "Success" : false, "Message": "failed to updateUser"}`
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "success"}`
	var us updateStats
	// 反序列化
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	//成功状态
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}

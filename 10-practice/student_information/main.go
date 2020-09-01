package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 批量插入一组学生信息，id,name,age,score
func studentAdd(num int) map[int]map[string]interface{} {
	student := make(map[int]map[string]interface{}, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		value, ok := student[i]
		if !ok {
			value = make(map[string]interface{}, 4)
		}
		id := i + 1
		value["id"] = id
		value["name"] = fmt.Sprintf("student%d", id)
		value["age"] = rand.Intn(10) + 10
		value["score"] = rand.Intn(40) + 60
		student[i] = value
	}
	return student
}

func main() {
	students := studentAdd(10)
	for index, value := range students {
		fmt.Printf("Index %d information:id=%d,name=%s,age=%d,sorce=%d\n", index, value["id"], value["name"], value["age"], value["score"])
	}
}

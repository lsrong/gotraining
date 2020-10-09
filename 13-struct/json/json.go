package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Sex  string
}
type Class struct {
	Name     string
	Count    int
	Students []*Student
}

var JsonData = `{"Name":"Class name","Count":10,"Students":[{"Id":0,"Name":"stu0","Sex":"man"},{"Id":1,"Name":"stu1","Sex":"man"},{"Id":2,"Name":"stu2","Sex":"man"},{"Id":3,"Name":"stu3","Sex":"man"},{"Id":4,"Name":"stu4","Sex":"man"},{"Id":5,"Name":"stu5","Sex":"girl"},{"Id":6,"Name":"stu6","Sex":"girl"},{"Id":7,"Name":"stu7","Sex":"girl"},{"Id":8,"Name":"stu8","Sex":"girl"},{"Id":9,"Name":"stu9","Sex":"girl"}]}`

func NewStudent(name, sex string, id int) *Student {
	return &Student{
		Name: name,
		Sex:  sex,
		Id:   id,
	}
}

func main() {
	class := &Class{
		Name:  "Class name",
		Count: 10,
	}
	i := 0
	for i < class.Count/2 {
		class.Students = append(class.Students, NewStudent(fmt.Sprintf("stu%d", i), "man", i))
		i++
	}
	for i < class.Count {
		class.Students = append(class.Students, NewStudent(fmt.Sprintf("stu%d", i), "girl", i))
		i++
	}

	// 序列化成json字符串
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("Json marshal failure!")
		return
	}
	fmt.Println(string(data))

	// 反序列化json字符串
	var classUnMarshal = &Class{}
	err = json.Unmarshal([]byte(JsonData), classUnMarshal)
	if err != nil {
		fmt.Println("Json unmarshal failure!")
		return
	}
	for _, v := range classUnMarshal.Students {
		fmt.Printf("stu:%#v \n", v)
	}
}

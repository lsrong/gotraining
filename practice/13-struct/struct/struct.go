package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NewStudent 一般情况下结构体需要有初始化的构造函数
// 创建一个新的指针结构体
func NewStudent(id int, name string, age int) Student {
	return Student{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func main() {
	demoStd()
	demoStdJson()
}

// demoStd 演示结构体的基本使用
func demoStd() {
	// 第一种 var 关键字定义， 属性值会是对应类型的零值
	var std1 Student
	fmt.Printf("结构体属性默认值是类型的零值：%+v \n", std1)
	std1.Id = 1
	std1.Name = "刘盛荣"
	std1.Age = 30

	fmt.Printf("先定义后赋值：%+v \n", std1)

	// 声明并赋值的方式
	std2 := Student{Id: 2, Name: "杜俊", Age: 30}
	fmt.Printf("通过构造函数生成的：%+v \n", std2)

	// 构造函数生成方式：
	std3 := NewStudent(2, "唐兆阳", 30)
	fmt.Printf("通过构造函数生成的：%+v \n", std3)
}

// demoStdJson 演示结构体的json操作
func demoStdJson() {
	std := NewStudent(1, "lsrong", 30)
	stdJson, err := json.Marshal(std)
	if err != nil {
		log.Fatal("failed to call json.Marshal")
	}
	fmt.Printf("结构体序列化的字符串：%s \n", stdJson)
	initStd := `{"id":2,  "name":"demo-student", "age": 18}`
	demoStd := new(Student)
	err = json.Unmarshal([]byte(initStd), demoStd)
	if err != nil {
		log.Fatalf("failed to call json.Unmarshal, error: %s", err)
	}

	fmt.Printf("我是反序列化过来的信息：%v \n", initStd)
}

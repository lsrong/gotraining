package main

import (
	"encoding/json"
	"fmt"
	"github.com/learning_golang/13-struct/construct"
)

// 基本定义:⾯向对象是通过struct来实现的, struct是⽤户⾃定义的类型
type Student struct {
	Name string
	Sex  string
	Class
}
type Class struct {
	Name    string
	Address string
}

// 定义简单Struct 类型
func defineStruct() {
	// 初始化1
	var student01 Student
	student01.Name = "学生01"
	student01.Sex = "男"
	//student01.Class = "一年级1班"
	fmt.Printf("学生01:%v \n", student01)
	//fmt.Printf("Type of student01 var is %T \n", student01)

	// 初始化02
	student02 := Student{
		Name: "学生01",
		Sex:  "女",
		//Class: "二年级02班",
	}
	fmt.Printf("学生02:%v \n", student02)

	// 初始化 默认值 Struct
	var student03 Student
	fmt.Printf("学生03:%v \n", student03)
}

// 结构体类型的指针
func varStruct(student Student) {
	student.Name = "var Struct"
}
func pointStruct(student *Student) {
	student.Name = "point Struct"
}
func testPointStruct() {
	// 定义非指针结构体
	var student Student
	varStruct(student)
	fmt.Printf("after call var struct %v \n", student)
	pointStruct(&student)
	fmt.Printf("after call point struct %v \n", student)

	// 直接定义结构体指针类型,需要指定结构体的地址
	var pointStudent = &Student{}
	// 指针类型结构体不能在传值函数中使用
	//varStruct(*pointStudent)
	//fmt.Printf("after call var struct %v \n", pointStudent)
	pointStruct(pointStudent)
	fmt.Printf("after call point struct %v \n", pointStudent)

	// new 实例化结构体指针类型
	newStudent := new(Student)
	pointStruct(newStudent)
	fmt.Printf("after call point struct %v \n", newStudent)

	// &User{}和new(User)
	//本质上是⼀样的，都是返回⼀个
	//结构体的地址
}

// 结构体没有构造函数， 必要时需要⾃⼰实现
func constructUser() {
	name := "nickname"
	username := "username"
	password := "123456"
	user := construct.NewUser(name, username, password)
	fmt.Println("New user by construct function", user)
}

// 匿名属性的方式实现继承
func anonymousStruct() {
	var student Student = Student{
		Name: "stu01",
		Sex:  "man",
		Class: Class{
			"class one",
			"First floor",
		},
	}
	fmt.Println(student.Name)
	fmt.Println(student.Class)
}

// 可见性：⼤写表示可公开访问，⼩写表示私有
// tag是结构体的元信息，可以在运⾏的时候通过反射的机制读取出来
// Json 序列化，通过tag定义序列化的键
func jsonStruct() {
	type User struct {
		Name     string `json:"name",db:"name"`
		Username string `json:"username",db:"user_name"`
		Password string `json:"password",db:"pwd"`
	}
	// 序列化
	user := new(User)
	userJson, _ := json.Marshal(user)
	fmt.Println(string(userJson))

	// 反序列化
	initUserJson := `{"name":"jsonUser","username":"jsonUsername","password":"jsonPassword"}`
	var unMarshalUser = new(User)
	_ = json.Unmarshal([]byte(initUserJson), unMarshalUser)
	fmt.Println(unMarshalUser)
}

func main() {

	defineStruct()

	testPointStruct()

	constructUser()

	anonymousStruct()

	jsonStruct()
}

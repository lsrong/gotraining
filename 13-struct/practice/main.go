package main

import (
	"fmt"
	"github.com/learning_golang/13-struct/practice/student"
	"os"
	"time"
)

var (
	students []*student.Student
	choose   int
)

func handleChoose() {
	_, _ = fmt.Scanf("%d\n", &choose)
	switch choose {
	case 1:
		showStudent()
	case 2:
		addStudent()
	case 3:
		editStudent()
	case 4:
		deleteStudent()
	case 5:
		os.Exit(0)
	}
}
func showStudent() {
	if len(students) == 0 {
		fmt.Println("暂无学生")
		return
	}
	for i, v := range students {
		fmt.Printf("学生【%d】:%v \n", i, v)
	}
}
func addStudent() {
	newStu := getStudentInformation()
	for i, v := range students {
		if v.Name == newStu.Name {
			students[i] = newStu
			return
		}
	}
	students = append(students, newStu)
	fmt.Println("添加成功！")
}
func editStudent() {
	newStu := getStudentInformation()
	for i, v := range students {
		if v.Name == newStu.Name {
			students[i] = newStu
			fmt.Println("修改成功！")
			return
		}
	}
	fmt.Println("学生没有找到！")
}
func getStudentInformation() *student.Student {
	var (
		name  string
		sex   string
		scope int
		grade string
	)
	fmt.Println("请输入学生姓名")
	_, _ = fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学生性别")
	_, _ = fmt.Scanf("%s\n", &sex)
	fmt.Println("请输入学生分数")
	_, _ = fmt.Scanf("%d\n", &scope)
	fmt.Println("请输入学生班级")
	_, _ = fmt.Scanf("%s\n", &grade)

	return student.NewStudent(name, sex, scope, grade)
}
func deleteStudent() {
	var name string
	fmt.Println("请输入需要删除的学生")
	_, _ = fmt.Scanf("%s\n", &name)
	for i, v := range students {
		if v.Name == name {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("学生没有找到！")
}
func main() {
	for {
		fmt.Println("模拟学生管理，以下为学生相关操作选项：")
		fmt.Println("1.列出学生")
		fmt.Println("2.新增学生")
		fmt.Println("3.修改学生")
		fmt.Println("4.删除学生")
		fmt.Println("5.退出")
		handleChoose()
		time.Sleep(1 * time.Second)
	}
}

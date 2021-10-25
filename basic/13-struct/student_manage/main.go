package main

import (
	"fmt"
	"os"
)

var (
	StudentManage = &Manager{}
	choose        int
	name          string
)

func showMenu() {
	fmt.Println("模拟学生管理，以下为学生相关操作选项：")
	fmt.Println("1.列出学生")
	fmt.Println("2.新增学生")
	fmt.Println("3.修改学生")
	fmt.Println("4.删除学生")
	fmt.Println("5.退出")
}

func ScanStudent() *Student {
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

	return NewStudent(name, sex, scope, grade)
}

func main() {
	for {
		showMenu()
		_, _ = fmt.Scanf("%d\n", &choose)
		switch choose {
		case 1:
			StudentManage.ShowStudent()
		case 2:
			StudentManage.AddStudent(ScanStudent())
		case 3:
			StudentManage.EditStudent(ScanStudent())
		case 4:
			fmt.Println("请输入需要删除的学生")
			_, _ = fmt.Scanf("%s\n", &name)
			StudentManage.deleteStudent(name)
		case 5:
			os.Exit(0)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/learning_golang/practice/13-struct/execrise/student"
)

var (
	choose  int
	name    string
	manager = &student.Manager{}
)

// 简易的学生命令行界面管理的小程序
// 功能：
// 命令行输入学生信息，姓名，性别，分数，班级
// 学生信息的增 删 改 查
func main() {
	for {
		shouMenu()
		_, err := fmt.Scanf("%d\n", &choose)
		exitErr(err)
		switch choose {
		case 1:
			manager.Show()
		case 2:
			manager.Add(ScanStudent())
		case 3:
			manager.Edit(ScanStudent())
		case 4:
			fmt.Println("请输入要删除的学员名字")
			_, err := fmt.Scanf("%s\n", &name)
			exitErr(err)
			manager.Delete(name)
		case 5:
			os.Exit(0)
		}

	}
}

func shouMenu() {
	menu := `模拟学生管理，以下为学生相关操作选项：
1.列出学生
2.新增学生
3.修改学生
4.删除学生
5.退出`
	fmt.Println(menu)
}

func exitErr(e error) {
	if e != nil {
		log.Fatalf("异常情况， error: %v", e)
	}
}

func ScanStudent() *student.Student {
	var (
		name  string
		sex   string
		scope int
		grade string
		err   error
	)
	fmt.Println("请输入姓名")
	_, err = fmt.Scanf("%s\n", &name)
	exitErr(err)
	fmt.Println("请输入性别")
	_, err = fmt.Scanf("%s\n", &sex)
	exitErr(err)
	fmt.Println("请输入分数")
	_, err = fmt.Scanf("%d\n", &scope)
	exitErr(err)
	fmt.Println("请输入班级")
	_, err = fmt.Scanf("%s\n", &grade)
	exitErr(err)

	return student.NewStudent(name, sex, scope, grade)
}

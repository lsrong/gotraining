package main

import (
	"fmt"
	"log"
)

// Sample program to show the syntax of type assertions.
// 显示类型断言语法的示例程序

type user struct {
	id   int
	name string
}

//finder 表示查找用户的能力。
type finder interface {
	find(id int) (*user, error)
}

// userSVC 是一种与用户交互的服务。
type userSvc struct {
	host string
}

// find 使用指针语义实现 finder 接口。
func (*userSvc) find(id int) (*user, error) {
	return &user{id: id, name: "Li"}, nil
}

func main() {
	svc := userSvc{
		host: "localhost:3434",
	}
	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

// run 针对传入调用的具体数据执行查找操作。
func run(f finder) error {
	u, err := f.find(1)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// 理想情况下，finder 抽象将包含您关心的所有行为。
	// 但是，如果出于某种原因，您真的需要获取存储在接口中的具体值，该怎么办？

	// 不可以从存储在此接口变量中的具体 userSVC 类型指针访问“主机”字段
	// 接口类型的实现只知道数据有一个名为“find”的方法。
	// ../assertion.go:54:26: f.host undefined (type finder has no field or method host)
	//log.Println("queried", f.host)

	// 可以使用类型断言来获取存储在接口内的 userSVC 指针的副本。
	if svc, ok := f.(*userSvc); ok {
		log.Println("queried", svc.host)
	}

	return nil
}

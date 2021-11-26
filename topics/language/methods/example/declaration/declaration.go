package main

import "fmt"

// Sample program to show how to declare methods and how the Go
// compiler supports them.

// 声明方法以及编译如何支持方法. 所有类型都可支持定义方法

type user struct {
	name  string
	email string
}

// notify 定义值方法,接受者为值类型.
func (u user) notify() {
	fmt.Printf("Sending User Emain To %s<%s>\n", u.name, u.email)
}

// changeEmail 定义指针方法,接受者为指针类型.
func (u *user) changeEmail(email string) {
	u.email = email
}

// newUserVaule 返回值实例
func newUserVaule(name, email string) user {
	return user{name, email}
}

// newUserVaule 返回指针实例
func newUserPointer(name, email string) *user {
	return &user{name, email}
}

func main() {
	// user 类型的值可用于调用使用值和指针接收器声明的方法, go编译器实现值和指针接受值的自动装换:
	// T.method(receiver, parameters):
	// 如果为值方法:接受者为指针,转换为 *receiver
	// 如果为指针方法:接受者为值,转换为 &receiver, 前提接受者是可取值的
	li := user{"Li", "Li@email.com"}
	li.changeEmail("Li@gmail.com")
	li.notify()

	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@gmail.com")
	joan.notify()

	fmt.Println("********")
	// 不可寻址的值接受者,不能调用指针方法
	newUserVaule("Mi", "Mi@eamil.com").notify()
	//newUserVaule("Mi", "Mi@eamil.com").changeEmail("Mi@gmaim.com")
	/**
	./declaration.go:50:36: cannot call pointer method on newUserVaule("Mi", "Mi@eamil.com")
	./declaration.go:50:36: cannot take the address of newUserVaule("Mi", "Mi@eamil.com"), 自动寻址失败
	*/
	mi := newUserVaule("Mi", "Mi@eamil.com")
	mi.changeEmail("Mi@gmail.com")

	// 指针接受者,能调用指针方法和值方法,因为可以通过指针自动取到值,(*newUserPointer("Mi", "Mi@eamil.com"))
	newUserPointer("Mi", "Mi@eamil.com").notify()
	(*newUserPointer("Mi", "Mi@eamil.com")).notify() // 编译器会自动转换
	newUserPointer("Mi", "Mi@eamil.com").changeEmail("Mi@gmaim.com")

	// 不建议在值变量中调用指针方法, 下面不推荐使用
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}
	for _, u := range users {
		u.changeEmail("it@email.com")
	}

}

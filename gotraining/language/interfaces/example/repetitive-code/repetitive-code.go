package main

import "fmt"

// Sample program that could benefit from polymorphic behavior with interfaces.
// 可以从接口的多态行为中受益的示例程序.

type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "Going Go Programming"
	copy(b, s)

	return len(s), nil
}

type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	s := `{name:"bill", title:"developer"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// retrieve 行为在没有抽象之前需要对每个具体的数据结构定义一个方法
	_ = retrieveFile(f)
	_ = retrievePipe(p)
}

func retrieveFile(f file) error {
	data := make([]byte, 100)
	l, e := f.read(data)
	if e != nil {
		return e
	}

	fmt.Println(string(data[:l]))

	return nil
}

func retrievePipe(p pipe) error {
	data := make([]byte, 100)
	l, e := p.read(data)
	if e != nil {
		return e
	}

	fmt.Println(string(data[:l]))

	return nil
}

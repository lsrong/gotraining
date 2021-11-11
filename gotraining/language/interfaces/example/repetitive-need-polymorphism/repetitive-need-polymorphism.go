package main

import "fmt"

// Sample program to show how polymorphic behavior with interfaces.
// 示例程序展示接口的多态行为。

type reader interface {
	read(b []byte) (int, error)
}

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

	_ = retrieve(f)
	_ = retrieve(p)
}

// retrieve 接受一个reader接口,处理data.
func retrieve(r reader) error {
	data := make([]byte, 100)
	l, e := r.read(data)
	if e != nil {
		return e
	}

	fmt.Println(string(data[:l]))

	return nil
}

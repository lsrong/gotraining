package main

import "fmt"

type People struct {
	Name    string
	Sex     string
	Country string
}

// 定义一个方法
func (p *People) SetName(name string) {
	p.Name = name
}
func (p *People) GetName() {
	fmt.Println(p.Name)
}

func main() {
	people := People{}
	people.SetName("test")
	people.GetName()
}

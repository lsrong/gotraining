package _6_interface

import "fmt"

// 定义接口
type Employer interface {
	CalcSalary() float32
}

type Program struct {
	name  string
	base  float32
	extra float32
}

func NewProgram(name string, base float32, extra float32) *Program {
	return &Program{
		name:  name,
		base:  base,
		extra: extra,
	}
}

// 实现接口方法
func (p *Program) CalcSalary() float32 {
	return p.base + p.extra
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base float32, extra float32) *Sale {
	return &Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}

// 实现接口
func (s *Sale) CalcSalary() float32 {
	return s.base + s.extra
}

// 实现了对接口的高级编程，更加抽象
func CalcSalary(employers []Employer) float32 {
	var cost float32
	for _, employer := range employers {
		cost += employer.CalcSalary()
	}
	return cost
}

// 判断接口属于那个结构体
func Just(e Employer) {
	switch v := e.(type) {
	case *Sale:
		fmt.Printf("v is Sale, %v \n", v)
	case *Program:
		fmt.Printf("v is Program, %v \n", v)
	default:
		fmt.Println("Not support")
	}
}

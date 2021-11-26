package main

// This is an example of using type hierarchies with a OOP pattern.
// 这是将类型层次结构与 OOP 模式一起使用的示例。
// This is not something we want to do in Go. Go does not have the
// concept of sub-typing. All types are their own and the concepts of
// base and derived types do not exist in Go. This pattern does not
// provide a good design principle in a Go program.
// 这不是我们想要在 Go 中做的事情。
// Go 没有子类型的概念。
// 所有类型都是它们自己的，Go 中不存在基类和派生类型的概念。
// 这种模式在 Go 程序中没有提供好的设计原则。
// Go中没有继承的概念(没有子类型概念), 下面列出的面向对象的继承在go中不存在的.

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Printf("Animal [%s] speaking \n", a.Name)
}

type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("Dog [%s] speaking with a pack factor of %d\n", d.Name, d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf("Cat [%s] speaking with a climb factor of %d\n", c.Name, c.ClimbFactor)
}

func main() {
	// 面向属性分组的思维在go中是不推崇的。
	animals := []Animal{
		Dog{
			Animal: Animal{
				Name: "Fido",
			},
			PackFactor: 1,
		},
		Cat{
			Animal: Animal{
				Name: "Milo",
			},
			ClimbFactor: 1,
		},
	}
	// ./group_state.go:46:6: cannot use Dog{...} (type Dog) as type Animal in slice literal
	// ./group_state.go:52:6: cannot use Cat{...} (type Cat) as type Animal in slice literal

	for _, animal := range animals {
		animal.Speak()
	}
}

// =============================================================================

// NOTES:

// Smells:
// 	* The Animal type is providing an abstraction layer of reusable state.
// 	* The program never needs to create or solely use a value of type Animal.
// 	* The implementation of the Speak method for the Animal type is a generalization.
// 	* The Speak method for the Animal type is never going to be called.
// Animal 类型提供了一个可重用状态的抽象层。
// 该程序从不需要创建或单独使用 Animal 类型的值。
// Animal 类型的 Speak 方法的实现是一种概括。
// Animal 类型的 Speak 方法永远不会被调用。

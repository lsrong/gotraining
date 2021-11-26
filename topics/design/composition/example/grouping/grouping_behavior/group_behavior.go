package main

// This is an example of using composition and interfaces. This is
// something we want to do in Go. We will group common types by
// their behavior and not by their state. This pattern does
// provide a good design principle in a Go program.
// 这是使用组合和接口的示例。
// 这是我们想要在 Go 中做的事情。
// 我们将按行为而非状态对常见类型进行分组。
// 这种模式确实在 Go 程序中提供了一个很好的设计原则。

import "fmt"

// Speaker 提供行为抽象，来实现具有相同行为的分组归类。
type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("Dog [%s] speaking with a pack factor of %d\n", d.Name, d.PackFactor)
}

type Cat struct {
	Name        string
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf("Cat [%s] speaking with a climb factor of %d\n", c.Name, c.ClimbFactor)
}
func main() {
	// 面向属性分组的思维在go中是不推崇的。
	speakers := []Speaker{
		&Dog{
			Name:       "Fido",
			PackFactor: 1,
		},
		&Cat{
			Name:        "Milo",
			ClimbFactor: 1,
		},
	}
	// ./group_state.go:46:6: cannot use Dog{...} (type Dog) as type Animal in slice literal
	// ./group_state.go:52:6: cannot use Cat{...} (type Cat) as type Animal in slice literal

	for _, speaker := range speakers {
		speaker.Speak()
	}
}

// =============================================================================

// NOTES:

// Here are some guidelines around declaring types:
// 	* Declare types that represent something new or unique.
// 	* Validate that a value of any type is created or used on its own.
// 	* Embed types to reuse existing behaviors you need to satisfy.
// 	* Question types that are an alias or abstraction for an existing type.
// 	* Question types whose sole purpose is to share common state.
// 以下是有关声明类型的一些准则:
// * 声明代表新事物或独特事物的类型。
// * 验证任何类型的值是单独创建或使用的。
// * 嵌入类型以重用您需要满足的现有行为。
// * 问题类型是现有类型的别名或抽象。
// * 唯一目的是共享公共状态的问题类型。

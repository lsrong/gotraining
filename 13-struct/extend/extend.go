package main

import "fmt"

type Animal struct {
	Name string
	Area string
}

func (a *Animal) Talk() {
	fmt.Printf("I'm %s.I can talk. I come from %s", a.Name, a.Area)
}

type Mammal struct {
}

func (m *Mammal) Talk() {
	fmt.Println("I'm a Mammal, I can talk")
}

type Dog struct {
	Feet string
	*Animal
	*Mammal
}

// 重写父类方法
func (s *Dog) Talk() {
	fmt.Println("I'm a dog, I can talk")
}

func main() {
	// 多层继承
	dog := &Dog{
		Feet: "test",
		Animal: &Animal{
			Name: "Tom",
			Area: "USA",
		},
		Mammal: &Mammal{},
	}

	// 匿名冲突
	dog.Talk()
	dog.Mammal.Talk()
	dog.Animal.Talk()
}

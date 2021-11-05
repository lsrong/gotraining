package main

import "fmt"

type Animal struct {
	Name string
	mean bool
}

type AnimalSounder interface {
	MakeNoise()
}

type Dog struct {
	Animal
	BarkStrength int
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

func (a *Animal) PerformNoise(strength int, sound string) {
	if a.mean {
		strength *= 5
	}
	for i := 0; i < strength; i++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Println()
}

func (d *Dog) MakeNoise() {
	d.PerformNoise(d.BarkStrength, "BARK")
}

func (c *Cat) MakeNoise() {
	c.Basics.PerformNoise(c.MeowStrength, "MEOW")
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

func main() {
	myDog := &Dog{
		Animal{
			"Rover",
			false,
		},
		2,
	}
	myCat := &Cat{
		Basics: Animal{
			Name: "Julius",
			mean: true,
		},
		MeowStrength: 3,
	}
	MakeSomeNoise(myDog)
	MakeSomeNoise(myCat)
}

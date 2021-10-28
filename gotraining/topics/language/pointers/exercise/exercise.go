package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
	level int
}

func allocateLevel(u *user, level int) {
	u.level = level
}

func main() {
	u := user{
		name:  "blunt",
		email: "blunt@gmial.com",
		age:   47,
	}
	fmt.Printf("%+v \n", u)
	fmt.Printf("Name:%s, level is  %d \n", u.name, u.level)

	allocateLevel(&u, 8)

	fmt.Printf("%+v \n", u)
	fmt.Printf("Name:%s,  level id %d \n", u.name, u.level)

}

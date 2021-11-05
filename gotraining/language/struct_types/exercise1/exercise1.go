package main

import "fmt"

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
type user struct {
	name  string
	email string
	age   int16
}

func main() {
	// initialize with values and display each field.
	bill := user{
		name:  "Bill",
		email: "bill@gmail.com",
		age:   18,
	}

	fmt.Println("Name: ", bill.name)
	fmt.Println("Email: ", bill.email)
	fmt.Println("Age: ", bill.age)

	// anonymous struct type
	ed := struct {
		name  string
		email string
		age   int16
	}{
		name:  "Ed",
		email: "ed@gmail.com",
		age:   18,
	}
	fmt.Println("Name: ", ed.name)
	fmt.Println("Email: ", ed.email)
	fmt.Println("Age: ", ed.age)

}

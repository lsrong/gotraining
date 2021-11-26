// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.

package main

import (
	"fmt"
	"github.com/learning_golang/topics/language/exporting/exercise/exercise1/toy"
)

func main() {
	t := toy.New("Bat", 32)
	t.UpdateSold(28)
	t.UpdateOnHand(52)

	fmt.Printf("Name: %s \n", t.Name)
	fmt.Printf("Weight: %d \n", t.Weight)
	fmt.Printf("OnHand: %d \n", t.OnHand())
	fmt.Printf("Sold: %d \n", t.Sold())
}

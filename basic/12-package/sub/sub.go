package sub

import "fmt"

func Sub(a, b int) int {
	return a - b
}

func init() {
	fmt.Println("This sub init()")
}

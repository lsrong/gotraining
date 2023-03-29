package sub

import "fmt"

func init() {
	fmt.Println("this is sub init function!")
}

func Sub(a, b int) int {
	return a - b
}

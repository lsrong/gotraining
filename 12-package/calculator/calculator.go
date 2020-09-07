package calculator

import (
	"fmt"
	"github.com/learning_golang/12-package/sub"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return sub.Sub(a, b)
}

func init() {
	fmt.Println("This is calculator init")
}

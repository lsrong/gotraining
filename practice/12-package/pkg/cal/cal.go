package cal

import (
	"fmt"

	"github.com/learning_golang/practice/12-package/pkg/sub"
)

func init() {
	fmt.Println("this is cal init function!")
}

func Add(a, b int) int {
	return a + b
}

// Sub 这里引入一个子包（sub），意图是测试出init的引入顺序
func Sub(a, b int) int {
	return sub.Sub(a, b)
}

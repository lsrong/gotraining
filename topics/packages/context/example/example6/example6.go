package main

import (
	"context"
	"fmt"
	"sync"
)

// Sample program to show when a Context is canceled, all Contexts
// derived from it are also canceled.
// 示例程序显示一个上下文何时被取消，从它派生的所有上下文也被取消。
/**
The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc.
WithCancel、WithDeadline 和 WithTimeout 函数采用 Context（父）并返回派生的 Context（子）和 CancelFunc。

Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers.
调用 CancelFunc 会取消子项及其子项，删除父项对子项的引用，并停止任何关联的计时器。

Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires.
调用 CancelFunc 失败会泄漏子项及其子项，直到父项被取消或计时器触发。

The go vet tool checks that CancelFuncs are used on all control-flow paths.
go vet 工具检查是否在所有控制流路径上使用了 CancelFuncs。
*/

type myKey int

const key myKey = 0

func main() {
	// 生成可取消的上下文 context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 生成一组goroutine, 每个协程衍生新的子上下文,
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			// 派生出新的上下文
			c := context.WithValue(ctx, key, id)

			// 等待上下文取消
			<-c.Done()
			fmt.Println("Canceled: ", id)
		}(i)
	}

	// 取消上下文以及其派生出来的子上下文
	cancel()
	wg.Wait()
}

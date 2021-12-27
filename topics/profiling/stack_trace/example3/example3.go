package main

import "fmt"

type strace struct{}

func main() {
	sl := make([]string, 2, 4)
	var t strace
	t.Example(sl, "test", 10)
}

func (t *strace) Example(slice []string, s string, i int) {
	fmt.Printf("Receiver pointer: %p \n", t)
	panic("Went stack strace")
}

/**
Receiver pointer: 0x102526d30
panic: Went stack strace

goroutine 1 [running]:
main.(*strace).Example(0x102526d30, {0x14000096f28, 0x2, 0x4}, {0x102449466, 0x4}, 0xa)
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/stack_trace/example3/example3.go:15 +0x98
main.main()
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/stack_trace/example3/example3.go:10 +0x6c

Process finished with the exit code 2


(*trace).Example(0x102526d30, {0x14000096f28, 0x2, 0x4}, {0x102449466, 0x4}, 0xa),
// 1. 接受者为 *stace的调用（使用指针接收器的方法调用）
// 2. 值列表首先显示接受者的值

0x102526d30 : *strace
{0x14000096f28, 0x2, 0x4} : slice []string
{0x102449466, 0x4} : s string
0xa : i int
*/

package main

// 如何读取堆栈跟踪的示例程序。
func main() {
	example(true, false, true, 52)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) error {
	panic("Want stack strace")
}

/**
$ go run example2.go
panic: Want stack stace

goroutine 1 [running]:
main.example(0x1, 0x0, 0x1, 0x34)
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/example2/example2.go:10 +0x38
main.main()
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/example2/example2.go:5 +0x38
exit status 2

// 定义：
main.example(b1, b2, b3 bool, i uint8)

// 调用
main.example(true, false, true, 52)

// 参数值：
b1 Value:	0x1
b2 Value:	0x0
b3 Value:	0x1
i Value:	0x34
*/

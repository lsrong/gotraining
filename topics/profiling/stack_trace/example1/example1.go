package main

// 演示如何读取堆栈追踪的信息
//go:noinline
func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, s string, i int) error {
	// 模拟出现 stack strace, 堆栈追踪.
	panic("Want stack stace")
}

/**
$ go run example1.go
输出：
panic: Want stack strace

goroutine 1 [running]:
main.example({0x14000096f28, 0x2, 0x4}, {0x102f39d05, 0x5}, 0xa)
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/example1/example1.go:12 +0x38
main.main()
        /Users/lsrong/Work/Project/Go/src/github.com/lsrong/gotraining/topics/profiling/example1/example1.go:6 +0x64

说明：
// example函数声明：
main.example(slice []string, s string, i int)

// 调用
main.example(make([]string, 2, 4), "hello", 10)

// stace example调用参数值：十六进制
main.example({0x14000064f28, 0x2, 0x4}, {0x10437dd05, 0x5}, 0xa)
Slice Value: 	{0x14000064f28, 0x2, 0x4}			pointer : 0x14000064f28, length 0x2 ,  capacity: 0x4
String Value: 	{0x10437dd05, 0x5}					pointer : 0x10437dd05, length: 0x5
Int Value: 		0xa									Int: 0xa
*/

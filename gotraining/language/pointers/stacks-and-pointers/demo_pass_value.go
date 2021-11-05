package main

func main() {
	count := 10
	// 打印 "值", "指针"
	println("count: \t Value of [", count, "]\t Addr of [", &count, "]")

	// 传递 "值" count
	increment(count)

	println("count: \t Value of [", count, "]\t Addr of [", &count, "]")
}

// go:noinline
func increment(inc int) {
	// 递增 值传递 的参数 inc
	inc++
	println("count: \t Value of [", inc, "]\t Addr of [", &inc, "]")
}

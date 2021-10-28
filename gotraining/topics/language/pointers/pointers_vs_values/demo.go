package main

import "fmt"

type ByteSlice []byte

func Append(dst, src []byte) []byte {
	l := len(dst)
	if l+len(src) > cap(dst) {
		newSlice := make([]byte, (l+len(src))*2)
		copy(newSlice, dst)
		dst = newSlice
	}
	copy(dst[l:], src)

	return dst
}

// AppendByValue Append in values
func (slice ByteSlice) AppendByValue(data []byte) []byte {
	return Append(slice, data)
}

// AppendByPointer Append in pointers
func (slice *ByteSlice) AppendByPointer(data []byte) {
	p := *slice // 取指针真实数据副本做追加操作

	// 操作完成之后将结果赋回类型指针
	*slice = Append(p, data)
}

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	p := Append(*slice, data)
	*slice = p

	return len(data), nil
}

func main() {
	var b ByteSlice
	p := b.AppendByValue([]byte("hello world"))
	fmt.Printf("after call `AppendByValue` b is : %s \t result is :%s\n", b, p)

	// (&b) 取指符号调用的为指针方法.
	(&b).AppendByPointer([]byte("hello world"))
	fmt.Printf("after call `AppendByPointer` b is : %s\n", b)

	_, _ = fmt.Fprintf(&b, ", this is a good days!")
	fmt.Printf("after call Fprintf (io.Writer) is : %s\n", b)
}

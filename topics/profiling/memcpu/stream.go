package main

import (
	"bytes"
	"fmt"
	"io"
)

// 示例程序采用字节流并查找字节“elvis”，找到后将其替换为“Elvis”。
// 代码不能假设流中有任何换行符或其他定界符，并且代码必须假设流具有任意长度。
// 该解决方案不能有意义地缓冲到流的末尾，然后处理替换。

type demo struct {
	input  []byte
	output []byte
}

var data = []demo{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

func assembleStream(demos []demo) ([]byte, []byte) {
	var in, out []byte
	for _, d := range demos {
		in = append(in, d.input...)
		out = append(out, d.output...)
	}

	return in, out
}

func main() {
	var output bytes.Buffer
	in, out := assembleStream(data)
	find := []byte("elvis")
	repl := []byte("Elvis")

	// 第一种算法： algoOne
	fmt.Println("-------- Running Algorithm One")
	output.Reset()
	algoOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v \nInp:[%s] \nExp:[%s] \nGot:[%s] \n", matched == 0, in, out, output.Bytes())

	// 第二种算法： algoTwo
	fmt.Println("-------- Running Algorithm Two")
	output.Reset()
	algoTwo(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v \nInp:[%s] \nExp:[%s] \nGot:[%s] \n", matched == 0, in, out, output.Bytes())
}

// algoOne 实现算法1： 使用io.ReadFull() + 缓存块变量buf，长度和find保持一致，每次读取len(find)-1的字节流，如果有长度和值等于find时用repl替换。
// 否则写入不满足的字节数据，然后移动缓存变量到下一个位置字节，循环到字节不满足find，说明已经循环结束
func algoOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	in := bytes.NewBuffer(data)
	size := len(find)
	// 中间缓存块
	buf := make([]byte, size)
	end := size - 1

	// 第一次读缓存块
	if n, err := io.ReadFull(in, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}
	// 移动中间缓存块，匹配字节块.
	for {
		// 读取下一个字节, 代表已经io.EOF
		if _, err := io.ReadFull(in, buf[end:]); err != nil {
			output.Write(buf[:end])
			return
		}

		if bytes.Equal(find, buf) {
			// 相同匹配替换为指定为的repl
			output.Write(repl)

			// 再次初始化读物缓存块
			if n, err := io.ReadFull(in, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}
		// 不相同，则写入到output并移动到读取字节
		output.WriteByte(buf[0])

		// 下一个位置
		copy(buf, buf[1:])
	}
}

// algoTwo 第二种实现方式: 使用bytes.ReadByte(), 逐个读取字节，和find的对应位置匹配，记录相同字节的位置idx
// idx == len(find)则替换，不相同则继续匹配， 读取一次字节然后动态匹配。
func algoTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	// 数据缓存字节
	in := bytes.NewBuffer(data)
	size := len(find)
	idx := 0
	for {
		// 没有读到字节则退出
		b, e := in.ReadByte()
		if e != nil {
			return
		}

		// 判断位置是否匹配
		if b == find[idx] {
			idx++ // 位置+1

			// 与find的大小相同，则表示找到有相同的find,可以用repl替换
			if idx == size {
				idx = 0
				output.Write(repl)
			}

			continue
		}

		// 匹配部分相同字符
		if idx != 0 {
			// 写入部分匹配的数据
			output.Write(find[:idx])

			// b 相当于没有读取
			in.UnreadByte()

			//重置下标
			idx = 0

			continue
		}

		output.WriteByte(b)
		// 重置下标
		idx = 0
	}
}

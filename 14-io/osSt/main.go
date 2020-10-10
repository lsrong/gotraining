package main

import "os"

/**
os.Stdin：标准输⼊的⽂件实例，类型为*File
os.Stdout：标准输出的⽂件实例，类型为*File
os.Stderr：标准错误输出的⽂件实例，类型为*File
*/
func main() {
	var buf [100]byte
	// 读取输入
	_, _ = os.Stdin.Read(buf[:])

	// 输出终端
	_, _ = os.Stdout.WriteString(string(buf[:]))

}

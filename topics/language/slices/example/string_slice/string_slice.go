package main

import "fmt"
import "unicode/utf8"

// 关于字符串说明: https://go.dev/blog/strings
// 字符串包含任意字节,字符串相当于是字节数组.
// utf-8 Go source code is always UTF-8.

// Sample program to show how strings have a UTF-8 encoded byte array.

func main() {
	s := "世界 means world"

	var buf [utf8.UTFMax]byte

	//  字符串的range遍历为rune(UTF-8),而是不是byte
	for i, r := range s {
		// rune的字节数
		rl := utf8.RuneLen(r)

		// 拷贝指定数的字节
		si := i + rl
		copy(buf[:], s[i:si]) // Copy of rune from the string to our buffer.

		fmt.Printf("%2d,: %q; codepoint: %#6x; encoded bytes:%#v\n", i, r, r, buf[:rl])

	}
}

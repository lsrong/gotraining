package strreve

import (
	"unicode"
	"unicode/utf8"
)

// Go 中的字符串不限于 UTF-8，但 Go 对它有很好的支持，并且最自然地与 UTF-8 一起使用。
// 某些字符串转换在 UTF-8 中工作，字符串上的范围子句在 UTF-8 中工作。
// Go 在标准库中还有一个 Unicode 包，可以轻松识别此任务的组合字符。

// Bytes 无编码字节层面反转，utf-8字符串字节反转会出现乱码.
func Bytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-i-1]
	}
	return string(r)
}

// Runes 将其参数解释为 UTF-8 并忽略不构成有效 UTF-8 的字节。返回值为 UTF-8。
func Runes(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		// 跳过非法的字符
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

// CombiningChars 返回有效同时不拆开组合字符的UTF-8.
func CombiningChars(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(s))
	start := len(s)
	// 控制遍历步伐，找到下一个组合位置字符.
	for i := 0; i < len(p); {
		// 跳过错误字符编码
		if p[i] == utf8.RuneError {
			i++
			continue
		}

		j := i + 1
		// 下一个编码字符位置
		for j < len(p) && (unicode.Is(unicode.Me, p[j]) ||
			unicode.Is(unicode.Mc, p[j]) || unicode.Is(unicode.Mn, p[j])) {
			j++
		}

		// 从最后一位开始设置反转位置。
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}

		// 跳跃步伐
		i = j
	}

	return string(r[start:])
}

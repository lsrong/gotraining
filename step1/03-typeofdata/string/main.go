package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func testString() {
	/**
	单个字符(字母)，一般使用 byte 来保存，且使用单引号包裹
	*/

	/**
	1.格式化输出:
	%%	%字面量
	%b	二进制整数值，基数为2，或者是一个科学记数法表示的指数为2的浮点数
	%c	该值对应的unicode字符
	%d	十进制数值，基数为10
	%e	科学记数法e表示的浮点或者复数
	%E	科学记数法E表示的浮点或者附属
	%f	标准计数法表示的浮点或者附属
	%o	8进制度
	%p	十六进制表示的一个地址值
	%s	输出字符串或字节数组
	%T	输出值的类型，注意int32和int是两种不同的类型，编译器不会自动转换，需要类型转换。
	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%t	单词true或false
	%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	%U	表示为Unicode格式：U+1234，等价于"U+%04X"
	*/

	/**
	1.Go的字符串是由单个字节连接起来的，即Go字符串是一串固定长度的字符连接起来的字符序列,内容初始化后不能修改
	一个中文3字节
	*/
	var str1 string
	str1 = "tests "
	str2 := "中文"

	fmt.Println(str1[0])
	fmt.Println(len(str2))
	fmt.Println(str1 + str2)
	/**
	修改字符
	*/
	// 字节数组[]byte类型 构建临时字符串
	var str string = "hello"
	strTemp := []byte(str)
	fmt.Println(strTemp)

	strTemp[0] = 'c'
	strChange := string(strTemp)
	fmt.Println(strChange)

	// 切片:str[1:], 从1开始到最后一位
	test := "test"
	fmt.Println("he" + test[2:])

	/**
	3.常用的string函数
	*/
	// len():长度(字节)
	lenStr1 := "test"
	lenStr2 := "中文"
	fmt.Println(len(lenStr1))
	fmt.Println(len(lenStr2))
	// 统一字符的UTF-8格式长度,字符长度
	fmt.Println(utf8.RuneCountInString(lenStr2))
	// 字符遍历
	for i := 0; i < len(lenStr1); i++ { //for
		fmt.Println(i, string(lenStr1[i]))
	}
	for i, ch := range lenStr2 { // for range
		fmt.Println(i, string(ch))
	}

	// string():字符串转换
	num := 1
	fmt.Printf("%T", num)
	fmt.Printf("%T", string(num))
	fmt.Println("")

	// 使用+能够连接字符串。但是该操作并不高效（因为字符串在Go中是基本类型，每次拼接都是拷贝了内存！）。
	// Go1.10提供了类似Java的StringBuilder机制来进行高效字符串连接
	strBuilder1 := "hello "
	strBuilder2 := "world"
	fmt.Println(splice(strBuilder1, strBuilder2))

}

// bytes.Buffer 拼接字符串
func splice(str1 string, str2 string) string {
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)

	return stringBuilder.String()
}

// 练习
func testStr() {
	var a string
	a = "hello"
	fmt.Println(a)

	// 赋值
	b := a
	fmt.Println(b)

	c := "hello"
	fmt.Println(c)

	// 不能重新定义
	//c := "test"

	// 占位符号： %s
	fmt.Printf("a=%s,b=%s,c=%s", a, b, c)

	// 万能(自动匹配数据类型)占位符号： %v
	fmt.Printf("a=%v,b=%v,c=%v", a, b, c)

	// 反引号 ``,原样输出
	var d = "d:\nworld"
	fmt.Println(d)
	f := `
Output as it is
`
	fmt.Print(f)

	// 长度：len
	g := "len"
	gLen := len(g)
	fmt.Printf("Length of g is %d", gLen)
}

// 字符串常用操作
func strOperation() {
	// 拼接：+ , fmt.Sprintf()
	f, g := "hello", "world"
	h := f + g
	fmt.Println(h)

	i := fmt.Sprintf("%s%s", f, g)
	fmt.Println(i)

	// 分隔: strings.Splice(string, flagStr) 返回数组
	ips := "192.168.1.11;192.168.1.66"
	ipSplice := strings.Split(ips, ";")
	fmt.Printf("first ip: %s \n", ipSplice[0])
	fmt.Printf("second ip: %s \n", ipSplice[1])

	// 包含：strings.Contains(str, flagStr)
	result := strings.Contains(ips, "192.168.1.66")
	fmt.Println(result)

	// 前缀strings.HasPrefix，后缀strings.HasSuffix
	str := "http://www.baidu.com,baidu"
	if strings.HasPrefix(str, "http") {
		fmt.Println("str is http url")
	} else {
		fmt.Println("str is not http url")
	}
	if strings.HasSuffix(str, "baidu.com") {
		fmt.Println("str is baidu url")
	} else {
		fmt.Println("str is not baidu url")
	}

	// 出现位置：首位：string.Index(str, flag),末尾strings.LastIndex(str,flag)
	index := strings.Index(str, "baidu")
	fmt.Printf("First index of baidu:%d \n", index)

	lastIndex := strings.LastIndex(str, "baidu")
	fmt.Printf("Last index of baidu:%d \n", lastIndex)

	// 数组转成字符
	var strArr = []string{"192.168.1.1", "192.168.1.2"}
	joinStr := strings.Join(strArr, ";")

	fmt.Println(joinStr)
}

func main() {
	// 基本定义
	testString()
	testStr()

	// 常用操作
	strOperation()
}

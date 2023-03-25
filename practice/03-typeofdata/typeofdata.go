package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	testBoolean()
	testNumeric()
	testString()
	testTime()
}

// testBoolean  布尔类型
func testBoolean() {
	fmt.Println("布尔类型")

	// 关键字 bool
	var a bool
	fmt.Printf("a = %t \n", a)

	// 赋值
	a = true
	fmt.Printf("a = %t \n", a)

	// !
	a = !a
	fmt.Printf("a = %t \n", a)

	a = true
	var b bool
	if a == true && b == true {
		fmt.Println("It is  true")
	} else {
		fmt.Println("It is false")
	}
	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not rigntx")
	}
	fmt.Printf("a = %t, b = %t \n", a, b)
}

// testNumeric 数值类型
func testNumeric() {
	fmt.Println("数值类型")
	/**
	1.整型:
	有符号:
	int     32位系统占4字节（与int32范围一样），64位系统占8个节（与int64范围一样）
	int8    占据1字节   范围 -128 ~ 127
	int16   占据2字节   范围 -2(15次方) ~ 2（15次方）-1
	int32   占据4字节   范围 -2(31次方) ~ 2（31次方）-1
	int64   占据8字节   范围 -2(63次方) ~ 2（63次方）-1
	rune	int32的别称

	无符号
	uint	32位系统占4字节（与uint32范围一样），64位系统占8字节（与uint64范围一样）
	uint8   占据1字节   范围 0 ~ 255
	uint16  占据2字节   范围 0 ~ 2（16次方）-1
	uint32  占据4字节   范围 0 ~ 2（32次方）-1
	uint64  占据8字节   范围 0 ~ 2（64次方）-1
	byte	uint8的别称
	*/
	var a int8
	a = 100
	fmt.Printf("a = %d \n", a)

	var b int32 = 100
	fmt.Printf("a + b = %d \n", int32(a)+b)

	/**
	2.浮点类型:
	float32 单精度  占据4字节   范围 -3.403E38 ~ 3.403E38    (math.MaxFloat32)
	float64 双精度  占据8字节   范围 -1.798E208 ~ 1.798E308  (math.MaxFloat64)
	*/
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	/**
	注意:精度缺失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1=",num1)		// -123.00009
	fmt.Println("num2=",num2)		// -123.0000901
	*/
	var num1 float64 = -123.0000901
	var num2 float64 = -123.000901
	isFloatEqual := func(f1, f2, p float64) bool {
		return math.Abs(f1-f2) < p
	}
	fmt.Println("num1=", num1) // -123.00009
	fmt.Println("num2=", num2) // -123.0000901
	fmt.Println(isFloatEqual(num1, num2, 0.0001))

	/**
	3.NaN
	*/
	nan := math.NaN()
	fmt.Println(math.IsNaN(nan))
}

// testString 字符串类型
func testString() {
	// Go的字符串是由单个字节连接起来的，即Go字符串是一串固定长度的字符连接起来的字符序列,内容初始化后不能修改
	// 一个中文3字节
	str1, str2 := "hello", "世界"
	fmt.Println(str1[0])
	fmt.Println(len(str2))
	fmt.Println(str1 + str2)

	// 修改字符
	var str string = "hello"
	strTemp := []byte(str)
	fmt.Println(strTemp)
	strTemp[3] = 'c'
	strChange := string(strTemp)
	fmt.Println(strChange) //  helco

	// 字符串切片
	test := "test"
	fmt.Println("he" + test[2:]) // hest

	// 字符创遍历
	lenStr1 := "test"
	lenStr2 := "中文"
	// 统一UTF-8字符长度
	fmt.Println(utf8.RuneCountInString(lenStr2))
	for i := 0; i < len(lenStr1); i++ {
		fmt.Println(i, string(lenStr1[i]))
	}
	// for range 中文遍历
	for i, ch := range lenStr2 {
		fmt.Println(i, string(ch))
	}

	// bytes.Buffer 拼接字符串
	s1, s2 := "hello ", "world"
	splice := func(s1, s2 string) string {
		var sb bytes.Buffer
		sb.WriteString(s1)
		sb.WriteString(s2)
		return sb.String()
	}
	fmt.Println(splice(s1, s2))

	// 常用操作
	// 拼接
	a, b := "hello ", "世界"
	c := a + b
	d := fmt.Sprintf("%s%s", a, b)
	fmt.Println(c, d)

	// 分隔： strings.Splice
	ips := "192.168.1.11;192.168.1.66"
	ipss := strings.Split(ips, ";")
	fmt.Printf("first ip: %s \n", ipss[0])
	fmt.Printf("second ip: %s \n", ipss[1])

	// 包含：strings.Contains
	ret := strings.Contains(ips, "192.168.1.66")
	fmt.Println(ret)

	// 前缀 strings.HasPrefix, 后缀 strings.HasSuffix
	str = "https://www.baidu.com,baidu"
	if strings.HasPrefix(str, "https://") {
		fmt.Println("str is https url")
	} else {
		fmt.Println("str is not https url")
	}
	if strings.HasSuffix(str, "baidu.com") {
		fmt.Println("str is baidu url")
	} else {
		fmt.Println("str is not baidu url")
	}

	// 出现的位置 strings.Index, strings.LastIndex
	firstIndex := strings.Index(str, "baidu")
	fmt.Printf("Fisrt index of baidu: %d \n", firstIndex)
	lastIndex := strings.LastIndex(str, "baidu")
	fmt.Printf("Second index of baidu: %d \n", lastIndex)

	// 字符串切片装成字符
	var strArr = []string{"192.168.1.1", "192.168.1.2"}
	joinStr := strings.Join(strArr, ";")
	fmt.Println(joinStr)
}

// testTime 标准包 - 时间
func testTime() {
	now := time.Now()
	fmt.Printf("当前时间： %s \n", now)

	// 年月日时分秒， time.Now().Year() ...
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 使用Printf格式化输出时间
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)

	// 当前时间戳
	timestamp := now.Unix()
	fmt.Printf("Timestamp is %d \n", timestamp)

	// 时间戳转换时间格式
	timestampToDateTime := func(timestamp int64) {
		timeObj := time.Unix(timestamp, 0)
		// 格式化时间
		fmt.Println(timeObj.Format("2006/01/02 15:04:05"))
	}
	timestampToDateTime(timestamp)

	// 解析字符串时间日期
	datetime := "2023-03-24 20:14:30"
	timeObj, _ := time.Parse("2006-01-02 15:04:05", datetime)
	fmt.Println(timeObj)

	// 定时任务
	ticker := time.Tick(time.Second)
	// 每秒执行一段代码块
	for i := range ticker {
		fmt.Printf("定时任务: %v \n", i)
	}

}

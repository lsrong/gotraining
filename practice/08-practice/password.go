package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	h bool
	l int
	t string
)

const (
	Number  string = "0123456789"
	English string = "qweryuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	flagParams()
	pwd := createPassword()
	fmt.Printf("生成密码：%s \n", pwd)
}

// flagParams 处理命令行参数
func flagParams() {
	// 处理命令行输入，h - 帮助文档， l-密码长度， t-密码类型
	flag.BoolVar(&h, "h", false, "thie is help")
	flag.IntVar(&l, "l", 16, "-l 密码长度")
	flag.StringVar(&t, "t", "n", "-t 密码类型: n-[0,9] e-[a-zA-z] a-[0-3a-zA-Z]")
	flag.Parse()
	if h {
		flag.Usage()
	}
}

// createPassword 生产密码字符串
func createPassword() string {
	var (
		pwd   = make([]byte, l, l)
		chars = Number
	)
	switch t {
	case "e":
		chars = English
	case "a":
		chars = fmt.Sprintf("%s%s", Number, English)
	}
	for i := 0; i < l; i++ {
		index := rand.Intn(len(chars))
		pwd[i] = chars[index]
	}

	return string(pwd)
}

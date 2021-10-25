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
	NUMBER  string = "0123456789"
	ENGLISH string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SPECIAL string = "!@#$%^&*()~{}[]"
)

// 获取命令行输入参数
func flagParams() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.IntVar(&l, "l", 16, "-l 密码长度")
	flag.StringVar(&t, "t", "n",
		"-t 密码类型：n:[0-9]；e:[a-zA-Z]；m:[0-9a-zA-Z]；a:[0-9a-zA-Z]+特殊字符")
	flag.Parse()
	if h {
		flag.Usage()
	}
}

func createPassword() string {
	var (
		password   = make([]byte, l, l)
		characters = NUMBER
	)
	switch t {
	case "e":
		characters = ENGLISH
	case "m":
		characters = fmt.Sprintf("%s%s", NUMBER, ENGLISH)
	case "a":
		characters = fmt.Sprintf("%s%s%s", NUMBER, ENGLISH, SPECIAL)
	}
	for i := 0; i < l; i++ {
		index := rand.Intn(len(characters))
		password[i] = characters[index]
	}

	return string(password)
}

func main() {
	rand.Seed(time.Now().Unix())
	flagParams()
	password := createPassword()
	fmt.Printf("生成的密码为：%s", password)
}

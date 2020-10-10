package main

import (
	"flag"
	"fmt"
)

var (
	Dns       string
	IsNetwork bool
	Level     int
)

// 结合init进行初始化输入参数

func init() {
	flag.StringVar(&Dns, "d", "localhost:3306", "Dns information xxx")
	flag.BoolVar(&IsNetwork, "n", false, "network option")
	flag.IntVar(&Level, "l", 1, "level of xxxx")
	flag.Parse()
}

func main() {
	fmt.Printf("Dns:%s\n", Dns)
	fmt.Printf("IsNetwork:%t\n", IsNetwork)
	fmt.Printf("Level:%d\n", Level)
}

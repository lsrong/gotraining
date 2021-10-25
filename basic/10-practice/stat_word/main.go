package main

import (
	"flag"
	"fmt"
	"strings"
)

var words string

func flagWords() {
	flag.StringVar(&words, "w", "测试字符串", "-w 输入一串字符")
	flag.Parse()
	//fmt.Println(words)
}

func statWords() map[string]int {
	stat := make(map[string]int, 128)
	words := strings.Split(words, "")
	//fmt.Println(words)
	for _, v := range words {
		count, ok := stat[v]
		if ok {
			stat[v] = count + 1
		} else {
			stat[v] = 1
		}
	}
	return stat
}

func main() {
	flagWords()
	stat := statWords()
	for word, count := range stat {
		fmt.Printf("Number of  %s is:%d  \n", word, count)
	}
}

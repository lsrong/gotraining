package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var words string

func main() {
	// 运行 flagword 示例
	flagWord()
	stats := statWords()
	for i, v := range stats {
		fmt.Printf("【%s】的重复数为 %d \n", i, v)
	}

	// 测试学生信息
	num := 10
	stds := studentAdd(num)
	keys := []int{}
	for i := range stds {
		// fmt.Printf("学生信息： id=%d, name=%s, age=%d \n", std["id"], std["name"], std["age"])
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, key := range keys {
		std := stds[key]
		fmt.Printf("学生信息： id=%d, name=%s, age=%d \n", std["id"], std["name"], std["age"])
	}
}

// flagWord 参数
func flagWord() {
	flag.StringVar(&words, "w", "hello world", "-w 输入一行字符")
	flag.Parse()
}

// statWords 统计每个字出现的次数
func statWords() map[string]int {
	stats := make(map[string]int, 32)
	statsWords := strings.Split(words, "")
	for _, v := range statsWords {
		num, ok := stats[v]
		if ok {
			stats[v] = num + 1
		} else {
			stats[v] = 1
		}
	}
	return stats
}

// studentAdd 批量插入一组学生信息
func studentAdd(num int) map[int]map[string]interface{} {
	students := make(map[int]map[string]interface{}, num)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= num; i++ {
		student := map[string]interface{}{
			"id":   i,
			"name": fmt.Sprintf("std_%d", i),
			"age":  10 + rand.Intn(10),
		}
		students[i] = student
	}
	return students
}

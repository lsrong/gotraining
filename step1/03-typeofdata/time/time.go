package main

import (
	"fmt"
	"time"
)

// 基本使用
func testTime() {
	// 当前时间
	now := time.Now()
	fmt.Printf("current time:%s \n", now)

	// 年，月，日，小时，分钟，秒数, time.Now().Year()......
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)

	// 获取当前时间戳
	timestamp := time.Now().Unix()
	fmt.Printf("Timestemp is %d \n", timestamp)
}

// 时间戳转换成时间
func timestampToDateTime(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj.Format("2006/01/02 15:04:05"))
}

// 解析字符串时间日期
func DateTimeParse() {
	datetime := "2020-08-20 23:05:26"
	timeObj, _ := time.Parse("2006-01-02 15:04:05", datetime)
	fmt.Println(timeObj)
}

// 编写定时任务
func testTicker() {
	// 时间管道
	ticker := time.Tick(time.Second)

	// 循环
	for i := range ticker {
		fmt.Println(i)
	}
}

// 定时任务处理函数
func processTask() {
	fmt.Println("定时任务处理程序")
}

// 格式化输出时间格式
func formatTime() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

// 使用Printf()输出格式化时间
func formatTimeByPrintf() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)
}

// 秒常量
func testTimeConst() {
	// 纳秒
	fmt.Println(time.Nanosecond)
	// 微妙
	fmt.Println(time.Microsecond)
	// 毫秒
	fmt.Println(time.Millisecond)
	// 秒
	fmt.Println(time.Second)
}

func main() {
	// 记录程序执行时间，精确到微妙
	start := time.Now().UnixNano()

	testTime()
	timestampToDateTime(time.Now().Unix() + 3600)
	DateTimeParse()
	//testTicker()
	formatTime()
	formatTimeByPrintf()
	testTimeConst()

	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("Code usedtime: %d us", cost)
}

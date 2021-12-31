package main

import (
	"log"
	"os"

	"github.com/learning_golang/topics/profiling/project/service"
)

// init 设置日志log参数
func init() {
	// 参数格式：显示日期时间，微妙，文件名
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	// 打印到标准输出
	log.SetOutput(os.Stdout)
}

// main 应用程序入口
func main() {
	service.Start()
}

package main

import (
	"expvar"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/learning_golang/topics/profiling/project/service"
)

// init 设置日志log参数
func init() {
	// 参数格式：显示日期时间，微妙，文件名
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	// 打印到标准输出
	log.SetOutput(os.Stdout)
}

func expvars() {
	// Add goroutine counts to the variable set.
	gr := expvar.NewInt("Goroutines")
	go func() {
		for _ = range time.Tick(time.Millisecond * 250) {
			gr.Set(int64(runtime.NumGoroutine()))
		}
	}()
}

// main 应用程序入口
func main() {
	expvars()
	service.Start()
}

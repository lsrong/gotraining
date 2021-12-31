package service

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/lsrong/kit/web"
)

// Start 启动服务，绑定到指定的端口并开始监听请求
func Start() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	api := web.NewServer(shutdown)
	// 定义服务路由
	api.Handle("POST", "", "/search", searchHandler)
	api.Handle("GET", "", "/search", searchHandler)
	api.Handle("GET", "", "/static/*filepath", staticHandler)
	api.Handle("GET", "", "/", indexHandler)

	// 服务参数配置：ip,端口，超时时间
	host := "localhost:4000"
	readTimeout := 10 * time.Second
	writeTimeout := 30 * time.Second
	idleTimeout := 60 * time.Second
	svc := http.Server{
		Addr:           host,
		Handler:        api,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    idleTimeout,
	}

	errCh := make(chan error)
	// goroutine 开启侦听服务
	go func() {
		log.Printf("Service Started, Listening on http://%s \n", host)
		errCh <- svc.ListenAndServe()
	}()

	// 监听是否结束服务，实现优雅关闭服务
	select {
	case <-shutdown:
		log.Println("Starting shutdown...")
		svc.Close()
	case err := <-errCh:
		log.Printf("Api Server Error: %v \n", err)
	}
}

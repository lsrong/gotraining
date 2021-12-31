package service

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Start 启动服务，绑定到指定的端口并开始监听请求
func Start() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	fs := http.FileServer(http.Dir("static"))
	http.StripPrefix("/static/", fs)

	// 定义服务路由
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", indexHandler)

	// 服务参数配置：ip,端口，超时时间
	host := "localhost:4000"
	readTimeout := 10 * time.Second
	writeTimeout := 60 * time.Second
	idleTimeout := 60 * time.Second
	svc := http.Server{
		Addr:           host,
		Handler:        http.DefaultServeMux,
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

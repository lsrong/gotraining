package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Sample program that implements a web request with a context that is
// used to timeout the request if it takes too long.
// 使用上下文实现 Web 请求的示例程序，该上下文用于在请求时间过长时使请求超时。

// Server frameworks that want to build on Context should provide implementations of Context to bridge between their packages and those that expect a Context parameter.
// 想要构建的服务器框架Context应该提供Context在其包和期望Context 参数的包之间架起桥梁的实现

// 用context实现在web请求中的简单超时机制
func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 设置query
	q := req.URL.Query()
	q.Set("q", "test")
	req.URL.RawQuery = q.Encode()

	// 生成带有超时时间的上下文
	duration := 500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// 生成带有上下文的请求句柄
	req = req.WithContext(ctx)

	// 执行web请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// 打印到终端
	io.Copy(os.Stdout, resp.Body)

}

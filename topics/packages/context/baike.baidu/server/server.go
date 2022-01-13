package main

// 实现的示例是一个 HTTP 服务器，它通过将查询“golang”转发到 百度百科接口 并呈现结果来处理
// 诸如 search?q=golang&timeout=1s 之类的 URL。
// timeout 参数告诉服务器在该持续时间过去后取消请求。

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/learning_golang/topics/packages/context/baike.baidu/baike"
	"github.com/learning_golang/topics/packages/context/baike.baidu/userip"
)

// resultsTemplate 结果模板
var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
	<h1>{{.Result.Title}}</h1>

	<div>{{.Result.Abstract}}</div>

	<p>链接: <a href='{{.Result.URL}}' target='blank'>{{.Result.URL}}</a></p>

  	<p>results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))

func main() {
	http.HandleFunc("/search", handleSearch)

	log.Println("Server Listening on: http://localhost:4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Printf("server error: %v \n", err)
		return
	}
}

// handleSearch 处理 /search?q=golang&timeout=1s 的请求, 如有timeout参数则生成具有超时期限的上下文
func handleSearch(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	// 处理timeout参数
	timeout, err := time.ParseDuration(r.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	// 校验查询参数
	q := r.FormValue("query")
	if q == "" {
		http.Error(w, "not query", http.StatusBadRequest)
		return
	}

	// 解析用户 IP 保存在ctx 上下文中
	userIP, err := userip.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// 执行搜索业务
	since := time.Now()
	result, err := baike.Search(ctx, q)
	elapsed := time.Since(since)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板响应html搜索结果
	if err := resultsTemplate.Execute(w, struct {
		Result           baike.Result
		Timeout, Elapsed time.Duration
	}{
		Result:  result,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}

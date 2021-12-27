package main

// 示例程序实现了一个简单的 Web 服务，它将允许我们探索如何查看核心

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/json", sendJSON)
	log.Println("Server started, Listening on: http://0.0.0.0:4000")
	http.ListenAndServe(":4000", nil)
}

func sendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		"Liu", "lsrong0414@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&u)
}

/**
// 运行程序：
$ go build -o godebug
$ ./godebug

// 使用hey工具模拟网络调用
$ hey -m POST -c 8 -n 1000000 "http://localhost:4000/json"

// 发出退出信号
$ ctrl + \

// 查看核心转储.
2021/12/27 14:29:40 Server started, Listening on: http://0.0.0.0:4000
^\SIGQUIT: quit
PC=0x10095aa30 m=0 sigcode=0

goroutine 19 [running]:
net/url.shouldEscape(0x2f, 0x1)
        /usr/local/go/src/net/url/url.go:135 +0x180 fp=0x140002bd740 sp=0x140002bd740 pc=0x10095aa30
net/url.escape({0x140003e21e5, 0x5}, 0x1)
        /usr/local/go/src/net/url/url.go:288 +0x68 fp=0x140002bd800 sp=0x140002bd740 pc=0x10095b4d8
net/url.(*URL).setPath(0x1400016d050, {0x140003e21e5, 0x5})
        /usr/local/go/src/net/url/url.go:682 +0x98 fp=0x140002bd850 sp=0x140002bd800 pc=0x10095ced8

......

lr      0x10095b4d8
sp      0x140002bd740
pc      0x10095aa30
fault   0x10095aa30


// 通过使用 GOTRACEBACK 环境变量运行程序来获得更大的故障转储。
$ GOTRACEBACK=crash ./godebug
*/

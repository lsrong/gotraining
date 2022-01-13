package main

//编写一个执行模拟数据库调用的 Web 处理程序，但如果调用时间过长，则会根据上下文超时。您还将状态保存到上下文中.

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type key int

const userIPKey key = 0

func main() {
	routes()

	log.Println("Server Listening on: http://localhost:4000")
	if err := http.ListenAndServe("0.0.0.0:4000", nil); err != nil {
		log.Printf("Listen Serve: %v \n", err)
		return
	}
}

// routes 定义服务处理函数
func routes() {
	http.HandleFunc("/user", findUser)
}

// findUser 利用上下文来处理超时和状态。
func findUser(rw http.ResponseWriter, r *http.Request) {
	// 生成有超时时间的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// 保存 ip 地址到上下文,之后的业务会用到这个参数
	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	// 使用goroutine查询数据库信息,将用户信息返回到channel中
	user := make(chan *User, 1)
	go func() {
		// 读取上下文中保存的ip,并记录
		if userIP, ok := ctx.Value(userIPKey).(string); ok {
			log.Printf("Start DB for IP: %s\n", userIP)
		}

		// 模拟查询数据
		user <- readDatabase()
		log.Println("DB goroutine terminated!")
	}()

	// 等待查询数据库接口,超时处理
	select {
	case u := <-user:
		respond(rw, u, http.StatusOK)
		log.Printf("Sent StatusOK.")
		return

	case <-ctx.Done():
		e := struct{ Error string }{ctx.Err().Error()}
		respond(rw, &e, http.StatusRequestTimeout)
		log.Printf("Sent StatusRequestTimeout.")

		return
	}
}

// readDatabase 模拟读取数据库信息,此过程要消耗一定时间
func readDatabase() *User {
	u := User{
		Name:  "LSRONG",
		Email: "lsrong0414@gmail.com",
	}

	time.Sleep(100 * time.Millisecond)
	//time.Sleep(10 * time.Millisecond)

	return &u
}

// respond 响应数据给调用者
func respond(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)
	_ = json.NewEncoder(rw).Encode(v)
}

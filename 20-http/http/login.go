package http

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	USERNAME = "admin"
	PASSWORD = "admin"
)

// 登录处理函数

func handleLogin(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if "GET" == method {
		loginTemple(w, r)
		return
	}
	if "POST" == method {
		checkLogin(w, r)
		return
	}

	_, _ = fmt.Fprintf(w, "Golang http service")
}

// 登录页面
func loginTemple(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/login.html")
	if err != nil {
		_, _ = fmt.Fprintf(w, "login html failure, err:%v\n", err)
		return
	}

	_ = t.Execute(w, nil)
}

// 登录操作
func checkLogin(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	if USERNAME != username || PASSWORD != password {
		_, _ = fmt.Fprintf(w, "用户名密码错误！！！")
		return
	}

	_, _ = fmt.Fprintf(w, "用户:%s 登录成功\n", r.FormValue("username"))
}

func Login() {
	// 首页
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Golang http service")
	})
	// 登录处理
	http.HandleFunc("/login", handleLogin)

	// 静态文件处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 监听
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Http service listen failed; err:%v \n", err)
	}
}

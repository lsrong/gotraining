package main

// 简单演示使用http,性能分析工具 pprof。

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/json", sendJSON)

	log.Println("Server listening on: http://localhost:4000")
	if err := http.ListenAndServe("localhost:4000", nil); err != nil {
		log.Fatal(err)
	}
}

func sendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		"LiuShengRong", "lsrong0414@gmail.com",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&u)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Server error: %v", err)
	}
}

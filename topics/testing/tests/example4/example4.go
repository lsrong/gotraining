package main

import (
	"log"
	"net/http"

	"github.com/learning_golang/topics/testing/tests/example4/handlers"
)

// 测试内部端点.
func main() {
	handlers.Routers()

	log.Println("Server started, listening on: 0.0.0.0:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

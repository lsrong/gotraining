package handlers

// Sample test to show how to test the execution of an internal endpoint.
// 展示如何测试内部端点执行的示例测试。

// 测试命令： go test -run TestSendJSON -race -cpu 8

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	Routers()
}

func TestSendJSON(t *testing.T) {
	uri := "/json"
	statusCode := 200

	// 给于需要测试的端点，模拟http请求响应。
	r := httptest.NewRequest("GET", uri, nil) // request
	w := httptest.NewRecorder()               // response
	http.DefaultServeMux.ServeHTTP(w, r)

	// 状态码断言
	if w.Code != statusCode {
		t.Fatalf("should receive statuc code [%d] for the response, Received[%d]", statusCode, w.Code)
	}

	// 响应断言
	var u struct {
		Name  string
		Email string
	}
	err := json.NewDecoder(w.Body).Decode(&u)
	if err != nil {
		t.Fatalf("Should be able to decode the response: %v", err)
	}

	if u.Name != "Ls" {
		t.Fatalf("Should have \"Ls\" in the response: %s.", u.Name)
	}
	if u.Email != "Ls@gmail.com" {
		t.Fatalf("Should have \"Ls@gmail.com\" in the response: %s.", u.Email)
	}

	t.Logf("the response: %+v.", u)
}

package example5

import (
	"net/http"
	"testing"
)

// 演示如何编写基本子单元表的示例测试。Sub Tests
// 运行命令：
// go test -v
// go test -run TestDownload/statusOk -v
// go test -run TestDownload/statusNotFound -v
// go test -run TestParallelize -v

func TestDownload(t *testing.T) {
	tts := []struct {
		Name       string
		URL        string
		StatusCode int
	}{
		{"statusOk", "http://localhost:8080", http.StatusOK},
		{"statusNotFound", "http://localhost:8080/api", http.StatusNotFound},
	}

	// 每个测试使用子测试来完成
	for id, test := range tts {
		// 子测试方法.
		ft := func(t *testing.T) {
			resp, err := http.Get(test.URL)
			if err != nil {
				t.Fatalf("Test: %d, Should be able to make Get call: %v", id, err)
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			if test.StatusCode != resp.StatusCode {
				t.Fatalf("Test: %d, Should receive a %d status code: %d", id, test.StatusCode, resp.StatusCode)
			}

			t.Logf("Test %d, receive a %d status code", id, test.StatusCode)
		}
		t.Run(test.Name, ft)
	}
}

// TestParallelize  t.Parallel()并发执行多个子测试。
/**
lsrong@lsrong-Mac example5 % go test -run TestParallelize -v
=== RUN   TestParallelize
=== RUN   TestParallelize/statusOk
=== PAUSE TestParallelize/statusOk
=== RUN   TestParallelize/statusNotFound
=== PAUSE TestParallelize/statusNotFound
=== CONT  TestParallelize/statusOk
=== CONT  TestParallelize/statusNotFound
=== CONT  TestParallelize/statusOk
    example5_test.go:76: Test 1, receive a 404 status code
=== CONT  TestParallelize/statusNotFound
    example5_test.go:76: Test 1, receive a 404 status code
--- PASS: TestParallelize (0.00s)
    --- PASS: TestParallelize/statusOk (0.05s)
    --- PASS: TestParallelize/statusNotFound (0.08s)
PASS
ok      github.com/learning_golang/topics/testing/tests/example5        0.189s
*/
func TestParallelize(t *testing.T) {
	tts := []struct {
		Name       string
		URL        string
		StatusCode int
	}{
		{"statusOk", "http://localhost:8080", http.StatusOK},
		{"statusNotFound", "http://localhost:8080/api", http.StatusNotFound},
	}

	// 每个测试使用子测试来完成
	for id, test := range tts {
		// 子测试方法.
		ft := func(t *testing.T) {
			t.Parallel() // 等待其他任务并发执行

			resp, err := http.Get(test.URL)
			if err != nil {
				t.Fatalf("Test: %d, Should be able to make Get call: %v", id, err)
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			if test.StatusCode != resp.StatusCode {
				t.Fatalf("Test: %d, Should receive a %d status code: %d", id, test.StatusCode, resp.StatusCode)
			}

			t.Logf("Test %d, receive a %d status code", id, test.StatusCode)
		}
		t.Run(test.Name, ft)
	}
}

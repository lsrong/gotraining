package example2

// Sample test to show how to write a basic unit table test. 展示如何编写基本单元表测试的示例测试。
// 批量测试的简单示例，同一个功能的不同状态。

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	// 批量测试用例。
	ts := []struct {
		url        string
		statusCode int
	}{
		{"http://localhost:8080", http.StatusOK},
		{"http://localhost:8080/api", http.StatusNotFound},
	}

	for id, test := range ts {
		resp, err := http.Get(test.url)
		if err != nil {
			t.Fatalf("Test: %d, Should be able to make Get call: %v", id, err)
		}

		t.Cleanup(func() {
			_ = resp.Body.Close()
		})

		if test.statusCode != resp.StatusCode {
			t.Fatalf("Test: %d, Should receive a %d status code: %d", id, test.statusCode, resp.StatusCode)
		}

		t.Logf("Test %d, receive a %d status code", id, test.statusCode)
	}

}

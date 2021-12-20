package example1

// Sample test to show how to write a basic unit test.  展示如何编写基本单元测试的示例测试。
// 基础的单元测试， 测试下载文件的功能

import (
	"net/http"
	"testing"
)

// TestDownload validates the http Get function can download content
func TestDownload(t *testing.T) {

	url, successStatusCode := "https://www.ardanlabs.com/blog/index.xml", http.StatusOK
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("http.Get() shuold be able to make the Get call: %v", err)
	}

	if resp.StatusCode != successStatusCode {
		t.Fatalf("Should receive a %d status code: %d", successStatusCode, resp.StatusCode)
	}

	t.Logf("should receive a %d status code", successStatusCode)

	t.Cleanup(func() {
		resp.Body.Close()
	})
}

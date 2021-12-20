package example3

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Sample test to show how to mock an HTTP GET call internally.
// 展示如何在内部模拟 HTTP GET 调用的示例测试。

// respXml 模拟XML文档格式响应
var respXML = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
	<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
    <item>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>
`

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"puDate"`
	Item        []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// mockServer 返回一个可用的模拟服务端程序
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/xml")
		_, _ = fmt.Fprintln(w, respXML)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 验证下载请求功能和反序列化响应内容。
func TestDownload(t *testing.T) {
	server := mockServer()
	defer server.Close()

	statusCode := http.StatusOK

	// 测试是否可以请求成功
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("shuold be able to make the Get call: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != statusCode {
		t.Fatalf("Should receive a %d status code: %d", statusCode, resp.StatusCode)
	}
	t.Logf("should receive a %d status code", statusCode)

	// 测试响应结果
	var d Document
	if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
		t.Fatalf("should be able to unmarshal the response: %v", err)
	}
	if len(d.Channel.Item) != 1 {
		t.Fatalf("Should have 1 item in the respXML: %d", len(d.Channel.Item))
	}

	t.Logf("Response have 1 item int respXML.")

}

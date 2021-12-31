package search

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

const cmEngine = "网易新闻"

// 接口列表
var cmFeeds = []string{
	"https://c.m.163.com/nc/article/headline/T1348647853363/0-100.html",
}

// CMResult 网易新闻结果
type CMResult map[string][]CMItem
type CMItem struct {
	PostId      string `json:"postid"`
	Title       string `json:"title"`
	Description string `json:"digest"`
	Link        string `json:"url"`
}

// CM 网易搜索引擎，通过一个互斥锁
type CM struct {
	m map[string]*sync.Mutex
}

// NewCM 生成网易搜索引擎
func NewCM() *CM {
	return &CM{
		m: make(map[string]*sync.Mutex),
	}
}

// Search 实现搜索接口，结果通过一个channel传递
func (c *CM) Search(keyword string, found chan<- []Result) {
	var results []Result
	for _, feed := range cmFeeds {
		res, err := c.match(feed, keyword)
		if err != nil {
			log.Printf("cm search error: %v", err)
			continue
		}

		results = append(results, res...)
	}

	found <- results
}

// match 匹配单个搜索接口的关键字
func (c *CM) match(url, keyword string) ([]Result, error) {
	mu, ok := c.m[url]
	if !ok {
		mu = &sync.Mutex{}
		c.m[url] = mu
	}

	var cmResult CMResult
	mu.Lock()
	defer mu.Unlock()

	v, ok := cache.Get(url)
	switch {
	case ok:
		cmResult = v.(CMResult)
	default:
		// 请求数据接口
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// 序列化响应数据
		if err := json.NewDecoder(resp.Body).Decode(&cmResult); err != nil {
			return nil, err
		}

		// 重置缓存
		cache.Set(url, cmResult, cacheExpiration)
		log.Println("reloaded cache ", url)
	}

	var results []Result
	// 从结果汇总匹配关键字
	for _, items := range cmResult {
		for _, item := range items {
			if strings.Contains(item.Title, keyword) || strings.Contains(item.Description, keyword) {
				results = append(results, Result{
					Engine:  cmEngine,
					Title:   item.Title,
					Content: item.Description,
					Link:    item.Link,
				})
			}
		}
	}

	return results, nil
}

package main

import (
	"fmt"
	"time"
)

// This program defines a type Feed with two methods: Count and Fetch. Create a
// new type CachingFeed that embeds *Feed but overrides the Fetch method.
// 该程序使用两种方法定义了一个类型 Feed：Count 和 Fetch。
// 创建一个嵌入 Feed 但覆盖 Fetch 方法的新类型 CachingFeed。

// The CachingFeed type should have a map of Documents to limit the number of
// calls to Feed.Fetch.
// CachingFeed 类型应该有一个 Documents 映射来限制对 Feed.Fetch 的调用次数。

// Document 文档数据模型.
type Document struct {
	Key   string
	Title string
}

// Feed 操作提取文档.
type Feed struct{}

// Count 文档总数.
func (f *Feed) Count() int {
	return 6
}

// Fetch 根据参数key查询文档.
func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)
	doc := Document{
		Key:   key,
		Title: fmt.Sprintf("Title for %s", key),
	}

	return doc, nil
}

// ==================================================

// CachingFeed 保留已经查询到的Documents的副本, 嵌入Feed获取Fetch和Count方法,"重写"让Fetch方法有缓存的功能
type CachingFeed struct {
	Docs map[string]Document
	*Feed
}

// NewCachingFeed initializes a CachingFeed for use.
func NewCachingFeed(f *Feed) *CachingFeed {
	//docs := make(map[string]Document)
	//return &CachingFeed{
	//	Docs: docs,
	//	Feed: f,
	//}

	return &CachingFeed{
		Docs: make(map[string]Document),
		Feed: f,
	}
}

// Fetch calls the embedded type's Fetch method if the key is not cached.
// 如果键未缓存，Fetch 将调用嵌入类型的 Fetch 方法。
func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.Docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.Docs[key] = doc

	return doc, nil
}

// ==================================================

// FetchCounter 定义一个行为接口.
type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

// process 处理流程
func process(fc FetchCounter) {
	keys := []string{"a", "a", "a", "b", "b", "b"}
	for _, k := range keys {
		doc, err := fc.Fetch(k)
		if err != nil {
			fmt.Printf("Could not fetch %s : %v", k, err)
			return
		}
		fmt.Printf("%s: %v\n", k, doc)
	}

	fmt.Printf("Has %d docements\n", fc.Count())
}

func main() {
	fmt.Println("Using Feed:")
	process(&Feed{})

	fmt.Println("Using CachingFeed:")
	cf := NewCachingFeed(&Feed{})
	process(cf)
}

package search

import (
	"html/template"
	"time"

	gc "github.com/patrickmn/go-cache"
)

const (
	cacheExpiration      = 10 * time.Minute
	cacheCleanupInterval = time.Hour
)

var cache = gc.New(cacheExpiration, cacheCleanupInterval)

// Options 用于执行搜索的选项，
type Options struct {
	Keyword string
	CM      bool
	Baidu   bool
}

// Result 查找到的结果
type Result struct {
	Engine  string
	Title   string
	Link    string
	Content string
}

// TitleHTML html 编码
func (r Result) TitleHTML() template.HTML {
	return template.HTML(r.Title)
}

// ContentHTML html 编码
func (r Result) ContentHTML() template.HTML {
	return template.HTML(r.Content)
}

// Searcher 搜索引擎接口，使用不同引擎来查找结果
type Searcher interface {
	Search(keyword string, found chan<- []Result)
}

// Submit 使用goroutine和channels并发执行搜索
func Submit(opt Options) (final []Result) {
	// 如果关键字为空不执行搜索
	if opt.Keyword == "" {
		return
	}
	var searchers []Searcher
	// 如果选择网易
	if opt.CM {
		searchers = append(searchers, NewCM())
	}

	results := make(chan []Result)
	// 并发执行搜索任务
	for _, searcher := range searchers {
		go searcher.Search(opt.Keyword, results)
	}

	// 等待接受goroutine搜索结果
	for i := 0; i < len(searchers); i++ {
		found := <-results
		final = append(final, found...)
	}
	return
}

package search

import "html/template"

type Options struct {
	Keyword string
	CM      bool
	Baidu   bool
}

type Result struct {
	Engine  string
	Title   string
	Link    string
	Content string
}

func (r Result) TitleHTML() template.HTML {
	return template.HTML(r.Title)
}

func (r Result) ContentHTML() template.HTML {
	return template.HTML(r.Content)
}

// Searcher 搜索接口，对接不同搜索引擎
type Searcher interface {
	Search(keyword string, found chan<- []Result)
}

func Submit(opt Options) (result []Result) {
	if opt.Keyword == "" {
		return
	}

	return
}

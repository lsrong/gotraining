package service

import (
	"context"
	"html/template"
	"net/http"

	"github.com/learning_golang/topics/profiling/project/search"
)

func staticHandler(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	fs := http.FileServer(http.Dir("static"))
	http.StripPrefix("/static/", fs).ServeHTTP(w, r)

	return nil
}

func indexHandler(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	fvs := formValues(r)

	return respondHtml(w, fvs)
}

func searchHandler(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	fvs := formValues(r)
	results := []search.Result{
		{
			Title:   "雷军的方向错了？小米12X预约量不到两千，不及小米12Pro的零头",
			Engine:  "百度资讯",
			Content: "12月28日，米粉们期待已久的小米新品发布会顺利召开，全新的小米12系列机型也正式亮相。小米总裁雷军在发布会上表示，从小米12系列开始，小米将全面对标苹果，因此小米12一共有三款机型，其中小米12和小米12X对标的是iPhone13，而小米12Pro则对标iPhone13Pro。",
			Link:    "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9644914547180506157%22%7D&n_type=-1&p_from=-1",
		},
	}
	return respondHtml(w, fvs, results...)
}

// formValues 获取提交参数
func formValues(r *http.Request) map[string]interface{} {
	fvs := make(map[string]interface{})
	// 关键字
	fvs["keyword"] = r.FormValue("keyword")

	// 网易新闻搜索选项，
	fvs["cm"] = ""
	if r.FormValue("cm") == "on" {
		fvs["cm"] = "checked"
	}

	// 百度
	fvs["bd"] = ""
	if r.FormValue("bd") == "on" {
		fvs["bd"] = "checked"
	}

	return fvs
}

func respondHtml(w http.ResponseWriter, params map[string]interface{}, results ...search.Result) error {
	if len(results) > 0 {
		rssVars := map[string]interface{}{"Items": results}
		rssHTML, err := execTemp(resultTemp, rssVars)
		if err != nil {
			return err
		}
		params["Results"] = template.HTML(rssHTML)
	}

	searchHTML, err := execTemp(searchTemp, params)
	if err != nil {
		return err
	}
	vars := map[string]interface{}{
		"LayoutContent": template.HTML(searchHTML),
	}
	layout, err := execTemp(layoutTemp, vars)
	if err != nil {
		return err
	}
	_, err = w.Write(layout)
	return err
}

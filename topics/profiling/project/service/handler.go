package service

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/learning_golang/topics/profiling/project/search"
)

// indexHandler 显示首页界面
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fvs, _ := formValues(r)

	respondHtml(w, fvs)
}

// searchHandler 接受搜索请求，响应查找结果页面
func searchHandler(w http.ResponseWriter, r *http.Request) {
	fvs, searchOpt := formValues(r)
	// 执行搜索
	results := search.Submit(searchOpt)

	respondHtml(w, fvs, results...)
}

// formValues 获取提交参数
func formValues(r *http.Request) (map[string]interface{}, search.Options) {
	fvs := make(map[string]interface{})
	var searchOpt search.Options
	// 关键字
	fvs["keyword"] = r.FormValue("keyword")
	searchOpt.Keyword = r.FormValue("keyword")
	// 网易新闻搜索选项，
	fvs["cm"] = ""
	if r.FormValue("cm") == "on" {
		fvs["cm"] = "checked"
		searchOpt.CM = true
	}

	// 百度
	fvs["bd"] = ""
	if r.FormValue("bd") == "on" {
		fvs["bd"] = "checked"
		searchOpt.Baidu = true
	}

	return fvs, searchOpt
}

// respondHtml 响应html页面数据
func respondHtml(w http.ResponseWriter, params map[string]interface{}, results ...search.Result) {
	if len(results) > 0 {
		rssVars := map[string]interface{}{"Items": results}
		rssHTML, err := execTemp(resultTemp, rssVars)
		if err != nil {
			_, _ = fmt.Fprintln(w, err)
			return
		}
		params["Results"] = template.HTML(rssHTML)
	}

	searchHTML, err := execTemp(searchTemp, params)
	if err != nil {
		_, _ = fmt.Fprintln(w, err)
		return
	}
	vars := map[string]interface{}{
		"LayoutContent": template.HTML(searchHTML),
	}
	layout, err := execTemp(layoutTemp, vars)
	if err != nil {
		_, _ = fmt.Fprintln(w, err)
		return
	}
	_, _ = w.Write(layout)
}

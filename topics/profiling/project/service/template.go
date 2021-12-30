package service

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

var views = make(map[string]*template.Template)

const (
	layoutTemp = "layout"
	resultTemp = "result"
	searchTemp = "search"
)

func init() {
	pwd, _ := os.Getwd()

	loadTemplate(layoutTemp, fmt.Sprintf("%s/views/basic-layout.html", pwd))
	loadTemplate(resultTemp, fmt.Sprintf("%s/views/results.html", pwd))
	loadTemplate(searchTemp, fmt.Sprintf("%s/views/search.html", pwd))
}

// loadTemplate 加载模板文件到全局的views模板缓存.
func loadTemplate(name, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("loading template ", err)
	}
	_, ok := views[name]
	if ok {
		return
	}
	t, err := template.New(name).Parse(string(data))
	if err != nil {
		log.Fatalln("New Template ", err)
	}

	views[name] = t
}

// execTemp 解析指定模板，返回解析的字符
func execTemp(name string, vars map[string]interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	t, ok := views[name]
	if !ok {
		return nil, fmt.Errorf("undefined template[%s]", name)
	}
	err := t.Execute(buf, vars)
	if err != nil {
		return nil, fmt.Errorf("processing template[%s], error: %v", name, err)
	}

	return buf.Bytes(), nil
}

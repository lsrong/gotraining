package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 返回json
func json(c *gin.Context) {
	ret := Result{
		Code:    0,
		Message: "ok, json render",
	}
	c.JSON(http.StatusOK, ret)
}

// 返回xml
func xml(c *gin.Context) {
	ret := Result{
		Code:    0,
		Message: "ok, render xml",
	}
	c.XML(http.StatusOK, ret)
}

// 返回yaml
func yaml(c *gin.Context) {
	ret := Result{
		Code:    0,
		Message: "ok, render xml",
	}
	c.YAML(http.StatusOK, ret)
}

// 返回html
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Golang",
	})
}

// 返回资源文件

func Render() {
	router := gin.Default()

	// Render json
	router.GET("/json", json)

	// Render xml
	router.GET("/xml", xml)

	// Render yaml
	router.GET("/yaml", yaml)

	// Render html
	router.LoadHTMLFiles("/Users/lsrong/Work/Project/Go/src/github.com/LearningGolang/23-gin/example/index.tmpl")
	router.GET("/index", index)

	// Render static
	router.Static("/static", "/Users/lsrong/Work/Project/Go/src/github.com/LearningGolang/23-gin/example/static")

	err := router.Run(":8888")
	if err != nil {
		fmt.Printf("Gin server error:%v \n", err.Error())
	}
}

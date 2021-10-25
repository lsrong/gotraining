package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const ROOT_PATH = "/Users/lsrong/Work/Project/Test/%s"

// 简单请求
func pingHandle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ping: " + ctx.Request.Host,
	})
}
func indexHandle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "index: " + ctx.Request.Host,
	})
}

// Query 参数
func queryHandle(ctx *gin.Context) {
	keyword := ctx.DefaultQuery("keyword", "test")
	name := ctx.Query("name")
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "成功",
		"keyword": keyword,
		"name":    name,
	})
}

// Path 参数
func pathHandle(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.Param("status")
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"id":      id,
		"status":  status,
	})
}

// post From 参数
func postFormHandle(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	ctx.JSON(200, gin.H{
		"code":     200,
		"message":  "ok",
		"username": username,
		"password": password,
	})
}

// 单文件上传
func uploadHandle(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)

	filepath := fmt.Sprintf(ROOT_PATH, file.Filename)
	err = ctx.SaveUploadedFile(file, filepath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("File[%s] upload!", file.Filename),
	})
}

// 多文件上传
func uploadMultiHandle(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	for index, file := range files {
		log.Println(index)
		log.Println(file.Filename)
		filepath := fmt.Sprintf(ROOT_PATH, file.Filename)
		err := ctx.SaveUploadedFile(file, filepath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    -1,
				"message": err.Error(),
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("%d files upload", len(files)),
	})
}

func Server() {
	router := gin.Default()
	// 首页
	router.GET("/", indexHandle)
	// Ping
	router.GET("/ping", pingHandle)

	// Query
	router.GET("/user", queryHandle)

	// Path info
	router.GET("/user/:id/status/:status", pathHandle)

	// From-data
	router.POST("/login", postFormHandle)

	// Upload
	router.POST("/upload", uploadHandle)

	// Upload Multi
	router.POST("/batch/upload", uploadMultiHandle)

	err := router.Run(":8888")
	if err != nil {
		fmt.Printf("Gin server run failed,err:%v \n", err)
	}
}

package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}

func checkLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}

func Group() {
	router := gin.Default()
	// 分组 v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", login)
		v1.POST("login", checkLogin)
		v1.GET("/user", user)
	}

	// 分组v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", login)
		v2.POST("login", checkLogin)
		v2.GET("/user", user)
	}

	err := router.Run(":8888")
	if err != nil {
		fmt.Printf("Gin group server failed,err:%s", err.Error())
	}
}

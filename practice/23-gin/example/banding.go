package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Json 绑定结构体
func bindingJson(c *gin.Context) {
	var login Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  "ok",
		"username": login.Username,
		"password": login.Password,
	})
}

// 绑定表单参数
func bindingForm(c *gin.Context) {
	var login Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  "ok",
		"username": login.Username,
		"password": login.Password,
	})
}

// Server
func Banding() {
	router := gin.Default()
	router.POST("/loginJson", bindingJson)
	router.POST("/loginForm", bindingForm)

	err := router.Run(":8888")
	if err != nil {
		fmt.Printf("Gin server failed, err, %v \n", err)
	}
}

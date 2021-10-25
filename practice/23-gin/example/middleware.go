package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		// ##### 执行处理之前
		c.Next() // 执行handle Func
		// #### 执行处理之后
		cost := time.Since(t)
		log.Printf("total cost time:%d", cost)
	}
}

func Middleware() {
	router := gin.New()
	router.Use(CostMiddleware())
	router.GET("", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
		})
	})
	err := router.Run(":8888")
	if err != nil {
		fmt.Printf("Gin server error:%v \n", err.Error())
	}
}

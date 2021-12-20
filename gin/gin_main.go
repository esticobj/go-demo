package main

import (
	"example.com/m/gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	// r.GET("/", func(context *gin.Context) {
	// 	// context.String(http.StatusOK, "Hello World")
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"msg": "123567水电费!",
	// 	})
	// })
	routers.InitArticleRouters(r)
	routers.InitUserRouters(r)
	//3.监听端口，默认8080
	r.Run(":8080")
}

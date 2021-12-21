package routers

import (
	"example.com/m/gin/controllers"
	"example.com/m/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouters(engine *gin.Engine) {
	group := engine.Group("user")
	group.Use(middlewares.PrintExecuteTime)
	controller := controllers.UserController{}
	group.GET("/add", controller.Add)
	group.GET("/edit", controller.Edit)
}

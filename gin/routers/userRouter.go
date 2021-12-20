package routers

import (
	"example.com/m/gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRouters(engine *gin.Engine) {
	group := engine.Group("user")
	controller := controllers.UserController{}
	group.GET("/add", controller.Add)
	group.GET("/edit", controller.Edit)
}

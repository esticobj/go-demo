package routers

import (
	"example.com/m/gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitArticleRouters(engine *gin.Engine) {
	group := engine.Group("article")
	controller := controllers.ArticleController{}
	group.GET("/add", controller.Add)
	group.GET("/edit", controller.Edit)
}

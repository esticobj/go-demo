package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (c *ArticleController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "article add...")
}

func (c *ArticleController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "article edit...")
}

package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (c *UserController) Add(ctx *gin.Context) {
	time.Sleep(time.Second)
	ctx.String(http.StatusOK, "user add...")
}

func (c *UserController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "user edit...")
}

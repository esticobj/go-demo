package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PrintExecuteTime(context *gin.Context) {
	start := time.Now().UnixMilli()
	context.Next()
	fmt.Printf("%v执行时长: %v\n", context.Request.RequestURI, (time.Now().UnixMilli() - start))
}

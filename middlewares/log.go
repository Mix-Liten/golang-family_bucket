package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context)  {

	fmt.Printf("UserAgent: %s\n", ctx.Request.UserAgent())
	ctx.Next()
}

package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(ctx *gin.Context)  {
	todoItems := [5]string{"a", "b", "c", "d"}
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todoItems,
	})
}
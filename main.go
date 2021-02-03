package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mix-liten/golang-family_bucket/api/todos"
	"github.com/mix-liten/golang-family_bucket/middlewares"
	"net/http"
	"strconv"
)

func hello(ctx *gin.Context) {
	name := ctx.Param("name")
	time, _ := strconv.Atoi(ctx.Query("time"))
	lastStr := "Nice to see you!"
	if time > 1 {
		lastStr = "Nice to see you again!"
	}
	ctx.String(http.StatusOK, "Hello %s, %s", name, lastStr)
}

func main()  {
	app := gin.Default()

	// /hello/name?time=3
	app.GET("/hello/:name", hello)

	todosApi := app.Group("/api/todos")
	todosApi.GET("/", middlewares.Log, todos.All)
	todosApi.GET("/:id")
	todosApi.PUT("/:id")
	todosApi.POST("/")
	todosApi.DELETE("/:id")

	app.Run()
}
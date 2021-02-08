package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mix-liten/golang-family_bucket/api/todos"
	"github.com/mix-liten/golang-family_bucket/libs/database"
	"github.com/mix-liten/golang-family_bucket/middlewares"
	"log"
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

func main() {
	defer func() {
		if err := database.DB.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	app := gin.Default()

	// /hello/name?time=3
	app.GET("/hello/:name", hello)

	todosApi := app.Group("/api/todos")
	todosApi.GET("/", middlewares.Log, todos.All)
	todosApi.GET("/:id", todos.One)
	todosApi.POST("/", todos.Create)
	todosApi.PUT("/:id", todos.Update)
	todosApi.DELETE("/:id", todos.Delete)

	app.Run()
}

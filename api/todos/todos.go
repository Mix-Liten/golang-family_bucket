package todos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mix-liten/golang-family_bucket/libs/database"
	_ "github.com/mix-liten/golang-family_bucket/libs/database"
	"github.com/mix-liten/golang-family_bucket/models"
	"net/http"
)

var (
	DB = database.DB
)

func All(ctx *gin.Context) {
	var todoItems []models.Todo
	// DB.Model(&todoItems).Where("Created_at = ?", time.Now()).Find(&todoItems)
	// DB.Raw("SELECT * FROM todos").Scan(&todoItems)
	//DB.Find(&todoItems)
	if err := DB.Find(&todoItems).Error; err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todoItems,
	})
}

func One(ctx *gin.Context) {
	id := ctx.Param("id")

	var todoItem models.Todo
	if err := DB.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todo": todoItem,
	})
}

type postCreate struct {
	Title string `json:"title" binding:"required"`
}

func Create(ctx *gin.Context) {
	var postData postCreate
	if err := ctx.Bind(&postData); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todoItem := models.Todo{
		Title: postData.Title,
	}

	if err := DB.Create(&todoItem).Error; err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": todoItem.ID,
	})
}

type postUpdate struct {
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done" binding:"required"`
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var postData postUpdate
	if err := ctx.Bind(&postData); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// if err := DB.Model(&todoItems).Where("id = ?", id).Update("title", postData.Title).Error

	var todoItem models.Todo
	if err := DB.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	todoItem.Title = postData.Title
	todoItem.Done = postData.Done
	if err := DB.Save(&todoItem).Error; err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	// if err := DB.Model(&todoItems).Where("id = ?", id).Update("title", postData.Title).Error

	var todoItem models.Todo
	if err := DB.Model(&todoItem).Where("id = ?", id).Delete(&todoItem).Error; err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

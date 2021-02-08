package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mix-liten/golang-family_bucket/models"
	"github.com/mosluce/go-toolkits"
	"github.com/mosluce/go-toolkits/database"
)

var DB *gorm.DB

func open() {
	var err error

	if gin.Mode() == gin.ReleaseMode {
		DB, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:  database.SQLITE,
			Filepath: "db.sqlite",
		})
	} else {
		DB, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:  database.SQLITE,
			Filepath: "db.sqlite",
		})

		DB.LogMode(true)
	}

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Todo{})
}

func init() {
	open()
}

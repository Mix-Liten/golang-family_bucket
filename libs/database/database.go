package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mix-liten/golang-family_bucket/models"
	"github.com/mosluce/go-toolkits"
	"github.com/mosluce/go-toolkits/database"
)

func Open() *gorm.DB {
	var db *gorm.DB
	var err error

	if gin.Mode() == gin.ReleaseMode {
		db, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:  database.SQLITE,
			Filepath: "db.sqlite",
		})
	} else {
		db, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:  database.SQLITE,
			Filepath: "db.sqlite",
		})

		db.LogMode(true)
	}

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})

	return db
}

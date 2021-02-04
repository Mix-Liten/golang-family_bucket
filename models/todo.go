package models

import (
	"time"
)

type Base struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//func (b Base) BeforeCreate(tx *gorm.DB) (err error) {
//	b.ID = xid.New().String()
//	return
//}

type Todo struct {
	Base
	Title string
	Done  bool
}

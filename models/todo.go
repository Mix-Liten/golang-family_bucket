package models

import (
	"github.com/rs/xid"
	"time"
)

type base struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (b base) BeforeCreate() (err error) {
	b.ID = xid.New().String()
	return
}

type Todo struct {
	base
	Title string
	Done  bool
}

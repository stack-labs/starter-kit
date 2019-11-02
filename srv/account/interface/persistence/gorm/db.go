package gorm

import (
	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {
	db := gorm.DB{}
	db.AutoMigrate()

	return &db
}

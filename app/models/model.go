package models

import (
	"github.com/jinzhu/gorm"
)

// Create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Post{})
	return db
}
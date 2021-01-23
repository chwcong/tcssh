package db

import "gorm.io/gorm"

func NewDB() *gorm.DB {
	return newSQlite()
}
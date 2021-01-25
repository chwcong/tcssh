package db

import (
	"gorm.io/gorm"
	"log"
	"sync"
	"tcssh/model"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func newDB() {
	once.Do(func() {
		DB = newSQlite()
	})
	return
}

func InitDB() {
	newDB()
	err := DB.AutoMigrate(&model.Dentry{})
	if err != nil {
		log.Println(err)
	}
	err = DB.AutoMigrate(&model.Node{})
	if err != nil {
		log.Println(err)
	}
}

package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
	"tcssh/util/config"
)

const dbFileName = "tcssh.conf"

func newSQlite() *gorm.DB {
	dsn := path.Join(config.WorkPath, dbFileName)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

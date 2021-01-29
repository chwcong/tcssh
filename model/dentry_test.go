package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
	"testing"
	"time"
)


func TestGetAllChildDentryByParentId(t *testing.T) {
	dir, _ := os.UserHomeDir()
	dsn := path.Join(dir,".tcssh/tcssh.conf")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	startTime := time.Now()
	result,err := GetAllChildDentryByParentId(db,1)
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
	fmt.Println(result)
	fmt.Println(err)
	//err =os.Remove(dsn)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
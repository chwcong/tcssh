package model

import (
	"tcssh/db"
	"time"
)

const GroupTableName =  "groups"

type Group struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentID  int       `gorm:"column:parent_id;index"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

func (g *Group) TableName() string {
	return GroupTableName
}

func GetGroupByParentID(parentId int) (groups []Group)  {
	db.DB.Table(GroupTableName).Where("parent_id = ?", parentId).Find(&groups)
	return
}

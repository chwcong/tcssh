package model

import (
	"gorm.io/gorm"
	"time"
)

const GroupTableName = "groups"

type Group struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentID    int       `gorm:"column:parent_id;index:parentId_idx;uniqueIndex:parent_group_name_idx"`
	Name        string    `gorm:"column:name;index:group_name_idx;uniqueIndex:parent_group_name_idx"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:create_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

func (g *Group) TableName() string {
	return GroupTableName
}

func GetGroupByParentID(db *gorm.DB, parentId int) (groups []Group) {
	db.Table(GroupTableName).Where("parent_id = ?", parentId).Find(&groups)
	return
}

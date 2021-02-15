package model

import (
	"gorm.io/gorm"
	"time"
)

const NodeTableName = "nodes"

type Node struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	DentryID     int       `gorm:"column:dentry_id;index:dentry_id"`
	Name        string    `gorm:"column:name"`
	Ip          string    `gorm:"column:ip"`
	UserName    string    `gorm:"column:user_name"`
	Password    string    `gorm:"column:password"`
	Port        int       `gorm:"column:port"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:create_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

func (n *Node) TableName() string {
	return NodeTableName
}

// DeleteNodeById delete one
func DeleteNodeById(db *gorm.DB,id int) error {
	return db.Table(NodeTableName).Delete(&Node{},id).Error
}

// DeleteNodeByIds batch delete
func DeleteNodesByIds(db *gorm.DB,ids []int) error {
	return db.Table(NodeTableName).Delete(&Node{},ids).Error
}

func CreateNode(db *gorm.DB,node *Node) (err error) {
	node.CreatedAt = time.Now()
	node.UpdatedAt = time.Now()
	err = db.Table(NodeTableName).Create(&node).Error
	return
}

func GetNodeOfCurrentIdByName(db *gorm.DB,id int,name string) (node *Node,err error) {
	err = db.Table(NodeTableName).Where("name = ? and dentry_id = ?",name, id).First(&node).Error
	return
}
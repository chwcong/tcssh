package model

import "time"

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

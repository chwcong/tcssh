package model

import "time"

const NodeTableName = "nodes"

type Node struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	GroupID     int       `gorm:"column:group_id;index:group_idx;uniqueIndex:group_node_name_idx"`
	Name        string    `gorm:"column:name;index:node_name_idx;uniqueIndex:group_node_name_idx"`
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

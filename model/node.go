package model

import "time"

type Node struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	GroupID   int       `gorm:"column:group_id;index"`
	Name      string    `gorm:"column:name"`
	Ip        string    `gorm:"column:ip"`
	UserName  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	Port      int       `gorm:"column:port"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

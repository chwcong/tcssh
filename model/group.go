package model

import (
	"time"
)

type Group struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentID  int       `gorm:"column:parent_id;index"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

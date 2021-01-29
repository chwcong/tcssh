package model

import (
	"gorm.io/gorm"
	"time"
)

const DentryTableName = "dentrys"

type Dentry struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentID    int       `gorm:"column:parent_id;index:parentId_idx;uniqueIndex:parent_dentry_idx"`
	Name        string    `gorm:"column:name;index:dentry_name_idx;uniqueIndex:parent_dentry_idx"`
	Description string    `gorm:"column:description"`
	NodeId      int       `gorm:"column:node_id"`
	Type        string    `gorm:"column:type"`
	CreatedAt   time.Time `gorm:"column:create_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

func (g *Dentry) TableName() string {
	return DentryTableName
}

func GetAllChildDentryByParentId(db *gorm.DB,id int) (dentrys []Dentry,err error)  {
	sql := `with recursive
    d as (
        select * from dentrys where dentrys.id = ?
        union all
        select  dentrys.* from d join dentrys on dentrys.parent_id = d.id
    )
SELECT * FROM d;`
	err = db.Table(DentryTableName).Raw(sql,id).Scan(&dentrys).Error
	return
}

// DeleteDentryByIds delete one
func DeleteDentryById(db *gorm.DB,id int) error {
	return db.Table(DentryTableName).Delete(&Dentry{},id).Error
}

// DeleteDentryByIds batch delete
func DeleteDentryByIds(db *gorm.DB,ids []int) error {
	return db.Table(DentryTableName).Delete(&Dentry{},ids).Error
}

func GetDentryByParentID(db *gorm.DB, parentId int) (dentrys []Dentry) {
	db.Table(DentryTableName).Where("parent_id = ?", parentId).Find(&dentrys)
	return
}

func GetDentryIdByNameAndParentID(db *gorm.DB,name string, parentId int) (dentry Dentry,err error) {
	err = db.Table(DentryTableName).Where("name = ? and parent_id = ?",name, parentId).First(&dentry).Error
	return
}

func CreateDentry(db *gorm.DB,dentry *Dentry) (err error) {
	dentry.CreatedAt = time.Now()
	dentry.UpdatedAt = time.Now()
	err = db.Table(DentryTableName).Create(&dentry).Error
	return
}

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// GroupAdmin ...
type GroupAdmin struct {
	ID        uint
	UserID    uint   `gorm:"column:user_id; not null"`
	Username  string `gorm:"-"`
	GroupID   uint   `gorm:"column:group_id; not null"`
	GroupName string `gorm:"-"`
}

// TableName ...
func (GroupAdmin) TableName() string {
	return "group_admin"
}

// Valid ...
func (GroupAdmin) Valid() error {
	return nil
}

//AddSQLConstrainsts ...
func (o *GroupAdmin) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddForeignKey("user_id", "system_user(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(user_id): %s", o.TableName(), e.Error())
		return
	}
	e = db.Model(o).AddForeignKey("group_id", "system_group(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(group_id): %s", o.TableName(), e.Error())
		return
	}

	e = db.Model(o).AddUniqueIndex("idx_system_user_group", "user_id", "group_id").Error
	if e != nil {
		e = fmt.Errorf("%s - AddIndex(): %s", o.TableName(), e.Error())
		return
	}
	return
}

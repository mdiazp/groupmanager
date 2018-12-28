package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// GroupADUser ...
type GroupADUser struct {
	ID        uint
	ADUser    string `gorm:"column:aduser; type:varchar(100); index; not null"`
	ADName    string `gorm:"column:adname; type:varchar(100); index; not null"`
	GroupID   uint   `gorm:"column:group_id; not null"`
	GroupName string `gorm:"-"`
}

// TableName ...
func (GroupADUser) TableName() string {
	return "group_aduser"
}

// Valid ...
func (GroupADUser) Valid() error {
	return nil
}

// Constrainsts
const (
	// ConstrainstGroupADUserUniqueIndex ...
	ConstrainstGroupADUserUniqueIndex = "uidx_aduser_group"
)

//AddSQLConstrainsts ...
func (o *GroupADUser) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddForeignKey("group_id", "system_group(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(group_id): %s", o.TableName(), e.Error())
		return
	}

	e = db.Model(o).AddUniqueIndex(ConstrainstGroupADUserUniqueIndex,
		"aduser", "group_id").Error
	if e != nil {
		e = fmt.Errorf("%s - AddIndex(): %s", o.TableName(), e.Error())
		return
	}
	return
}

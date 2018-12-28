package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Group ...
type Group struct {
	ID          uint
	Name        string `gorm:"type:varchar(20); unique_index; not null"`
	Description string `gorm:"type:varchar(500); not null"`
	Actived     bool   `gorm:"not null"`
}

// TableName ...
func (Group) TableName() string {
	return "system_group"
}

// Valid ...
func (o Group) Valid() error {
	lenName := len(o.Name)
	lenDescription := len(o.Description)

	e := ""
	if !(0 < lenName) {
		e = "Name can't be empty"
	}
	if !(lenName <= 20) {
		e = "Name's size > 20"
	}

	if !(0 < lenDescription) {
		e = "Description can't be empty"
	}
	if !(lenDescription <= 500) {
		e = "Description's size > 500"
	}
	if e != "" {
		return fmt.Errorf(e)
	}
	return nil
}

// AddSQLConstrainsts ...
func (o *Group) AddSQLConstrainsts(db *gorm.DB) error {
	return nil
}

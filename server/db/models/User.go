package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID       uint
	Provider string `gorm:"type:varchar(100);index;not null"`
	Username string `gorm:"type:varchar(100);unique_index;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Rol      string `gorm:"type:varchar(100);not null"`
	Enabled  bool   `gorm:"not null"`
}

// Valid ...
func (o User) Valid() error {
	return nil
}

// TableName ...
func (User) TableName() string {
	return "system_user"
}

// AddSQLConstrainsts ...
func (o *User) AddSQLConstrainsts(db *gorm.DB) error {
	return nil
}

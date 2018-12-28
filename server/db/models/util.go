package models

import (
	"github.com/jinzhu/gorm"
)

// GMModel ...
type GMModel interface {
	AddSQLConstrainsts(db *gorm.DB) error
}

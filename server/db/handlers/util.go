package handlers

import (
	"github.com/jinzhu/gorm"
)

// Paginator ...
type Paginator struct {
	Offset int
	Limit  int
}

// OrderBy ...
type OrderBy struct {
	By   string
	DESC bool
}

func orderByAndPaginator(db *gorm.DB, orderBy *OrderBy, pag *Paginator, tableName string) *gorm.DB {
	if orderBy != nil {
		kk := "asc"
		if orderBy.DESC {
			kk = "desc"
		}
		db = db.Order(tableName + "." + orderBy.By + " " + kk)
	}
	if pag != nil {
		db = db.Limit(pag.Limit).Offset(pag.Offset)
	}
	return db
}

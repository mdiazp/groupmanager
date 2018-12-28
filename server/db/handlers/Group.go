package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/gm/server/db/models"
)

// GroupHandler ...
type GroupHandler interface {
	CreateGroup(o *models.Group) error
	RetrieveGroupByID(id uint, o *models.Group) error
	UpdateGroup(id uint, o *models.Group) error
	DeleteGroup(id uint) error

	RetrieveGroupList(l *[]models.Group, filter *GroupFilter,
		orderBy *OrderBy, pag *Paginator) error
	CountGroups(filter *GroupFilter) (count int, e error)
}

// GroupFilter ...
type GroupFilter struct {
	NameSubstr *string
	Actived    *bool
	AdminID    *uint
	ADUser     *string
}

/////////////////////////////////////////////////////////////////////////////////////

func (h *handler) CreateGroup(o *models.Group) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveGroupByID(id uint, o *models.Group) error {
	return h.Where("id = ?", id).First(o).Error
}

func (h *handler) UpdateGroup(id uint, o *models.Group) error {
	return h.Save(o).Error
	// return h.Model(models.Group{}).Where("id=?", id).Update(o).Error
}

func (h *handler) DeleteGroup(id uint) error {
	return h.Delete(models.Group{ID: id}).Error
}

func (h *handler) RetrieveGroupList(l *[]models.Group, filter *GroupFilter,
	orderBy *OrderBy, pag *Paginator) error {
	db := makeGroupFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.Group{}).TableName())
	db = db.Find(l)
	fmt.Println("SQL Query = ", db.QueryExpr())
	e := db.Error
	return e
}

func (h *handler) CountGroups(filter *GroupFilter) (count int, e error) {
	e = makeGroupFilter(h.DB.Model(&models.Group{}), filter).Count(&count).Error
	return
}

func makeGroupFilter(db *gorm.DB, filter *GroupFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.NameSubstr != nil {
		db = db.Where("name ILIKE ?", "%"+*(filter.NameSubstr)+"%")
	}
	if filter.Actived != nil {
		db = db.Where("actived = ?", *(filter.Actived))
	}

	fmt.Println("AdminID filtering has to be doed by join to table group_admin")
	if filter.AdminID != nil {
		db = db.Where("id IN (?)",
			db.Model(&models.GroupAdmin{}).
				Select("group_id").Where("user_id = ?", *(filter.AdminID)).QueryExpr())
	}

	if filter.ADUser != nil {
		db = db.Where("id IN (?)",
			db.Model(&models.GroupADUser{}).
				Select("group_id").Where("aduser = ?", *(filter.ADUser)).QueryExpr())
	}

	return db
}

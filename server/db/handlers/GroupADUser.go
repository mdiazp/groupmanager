package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/gm/server/db/models"
)

// GroupADUserHandler ...
type GroupADUserHandler interface {
	CreateGroupADUser(o *models.GroupADUser) error
	RetrieveGroupADUser(groupID uint, aduser string) error
	DeleteGroupADUser(groupID uint, aduser string) error

	RetrieveGroupADUserList(filter *GroupADUserFilter, orderBy *OrderBy, pag *Paginator,
		l *[]models.GroupADUser) error
	CountGroupADUsers(filter *GroupADUserFilter) (count int, e error)
}

// GroupADUserFilter ...
type GroupADUserFilter struct {
	GroupID *uint
	ADUser  *string
}

/////////////////////////////////////////////////////////////////////////////////////
func makeGroupADUserFilter(db *gorm.DB, filter *GroupADUserFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.GroupID != nil {
		db = db.Where("group_aduser.group_id = ?", *(filter.GroupID))
	}
	if filter.ADUser != nil {
		db = db.Where("group_aduser.aduser = ?", *(filter.ADUser))
	}

	return db
}

func (h *handler) CreateGroupADUser(o *models.GroupADUser) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveGroupADUser(groupID uint, aduser string) error {
	o := &models.GroupADUser{}
	return h.Where("group_id = ? AND aduser = ?", groupID, aduser).First(o).Error
}

func (h *handler) DeleteGroupADUser(groupID uint, aduser string) error {
	db := h.DB.Where("group_id = ? AND aduser = ?", groupID, aduser)
	db = db.Delete(models.GroupADUser{})
	e := db.Error
	return e
}

func (h *handler) RetrieveGroupADUserList(filter *GroupADUserFilter, orderBy *OrderBy, pag *Paginator,
	l *[]models.GroupADUser) error {
	db := h.DB.Model(&models.GroupADUser{})
	db = makeGroupADUserFilter(db, filter)
	if orderBy == nil {
		orderBy = &OrderBy{By: "id", DESC: false}
	}
	db = orderByAndPaginator(db, orderBy, pag, (models.GroupADUser{}).TableName())

	db = db.
		Select(
			"group_aduser.id, group_aduser.aduser, group_aduser.adname, " +
				"group_aduser.group_id, system_group.name").
		Joins("left join system_group on group_aduser.group_id = system_group.id")

	rows, e := db.Rows()
	if e != nil {
		return e
	}
	defer rows.Close()

	for rows.Next() {
		o := models.GroupADUser{}
		e = rows.Scan(&o.ID, &o.ADUser, &o.ADName, &o.GroupID, &o.GroupName)
		if e != nil {
			return e
		}
		*l = append(*l, o)
	}

	return e
}

func (h *handler) CountGroupADUsers(filter *GroupADUserFilter) (count int, e error) {
	db := h.DB.Model(&models.GroupADUser{})
	db = makeGroupADUserFilter(db, filter)
	e = db.Count(&count).Error
	return
}

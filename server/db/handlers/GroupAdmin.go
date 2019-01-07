package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/gm/server/db/models"
)

// GroupAdminHandler ...
type GroupAdminHandler interface {
	CreateGroupAdmin(o *models.GroupAdmin) error
	RetrieveGroupAdmin(groupID uint, userID uint) error
	DeleteGroupAdmin(groupID uint, userID uint) error

	RetrieveGroupAdminList(filter *GroupAdminFilter,
		orderBy *OrderBy, pag *Paginator, l *[]models.GroupAdmin) error
	CountGroupAdmins(filter *GroupAdminFilter) (count int, e error)
}

// GroupAdminFilter ...
type GroupAdminFilter struct {
	GroupID      *uint
	GroupActived *bool
	UserID       *uint
}

//////////////////////////////////////////////////////////////////////////////////
func makeGroupAdminFilter(db *gorm.DB, filter *GroupAdminFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.GroupID != nil {
		db = db.Where("group_admin.group_id = ?", *(filter.GroupID))
	}
	if filter.UserID != nil {
		db = db.Where("group_admin.user_id = ?", *(filter.UserID))
	}
	if filter.GroupActived != nil {
		db = db.Where("system_group.actived = ?", *(filter.GroupActived))
	}

	return db
}

func addJoin(db *gorm.DB) *gorm.DB {
	return db.Joins("left join system_user on group_admin.user_id = system_user.id " +
		"left join system_group on group_admin.group_id = system_group.id")
}

func (h *handler) CreateGroupAdmin(o *models.GroupAdmin) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveGroupAdmin(groupID uint, userID uint) error {
	o := &models.GroupAdmin{}
	return h.Where("group_id = ? AND user_id = ?", groupID, userID).First(o).Error
}

func (h *handler) DeleteGroupAdmin(groupID uint, userID uint) error {
	db := h.DB.Where("group_id = ? AND user_id = ?", groupID, userID)
	db = db.Delete(models.GroupAdmin{})
	e := db.Error
	return e
}

func (h *handler) RetrieveGroupAdminList(filter *GroupAdminFilter,
	orderBy *OrderBy, pag *Paginator, l *[]models.GroupAdmin) error {

	db := h.DB.Model(&models.GroupAdmin{})
	db = makeGroupAdminFilter(db, filter)
	db = addJoin(db)
	if orderBy == nil {
		orderBy = &OrderBy{By: "id", DESC: false}
	}

	db = orderByAndPaginator(db, orderBy, pag, (models.GroupAdmin{}).TableName())

	db = db.
		Select(
			"group_admin.id, group_admin.user_id, group_admin.group_id, " +
				"system_user.username, system_group.name")

	rows, e := db.Rows()
	if e != nil {
		return e
	}
	defer rows.Close()

	for rows.Next() {
		o := models.GroupAdmin{}
		e = rows.Scan(&o.ID, &o.UserID, &o.GroupID, &o.Username, &o.GroupName)
		if e != nil {
			return e
		}
		*l = append(*l, o)
	}

	return e
}

func (h *handler) CountGroupAdmins(filter *GroupAdminFilter) (count int, e error) {
	db := h.DB.Model(&models.GroupAdmin{})
	db = makeGroupAdminFilter(db, filter)
	db = addJoin(db)
	db = db.Count(&count)
	e = db.Error
	return
}

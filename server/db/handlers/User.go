package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/gm/server/db/models"
)

// UserHandler ...
type UserHandler interface {
	CreateUser(o *models.User) error
	RetrieveUserByID(id uint, o *models.User) error
	RetrieveUserByUsername(username string, o *models.User) error
	UpdateUser(id uint, o *models.User) error
	DeleteUser(id uint) error

	RetrieveUserList(l *[]models.User, filter *UserFilter,
		orderBy *OrderBy, pag *Paginator) error
	CountUsers(filter *UserFilter) (count int, e error)
}

// UserFilter ...
type UserFilter struct {
	UsernamePrefix  *string
	Provider        *string
	Rol             *string
	NameSubstr      *string
	Enabled         *bool
	GroupWhichAdmin *int
}

/////////////////////////////////////////////////////////////////////////////////////

func (h *handler) CreateUser(o *models.User) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveUserByID(id uint, o *models.User) error {
	return h.Where("id = ?", id).First(o).Error
}

func (h *handler) RetrieveUserByUsername(username string, o *models.User) error {
	return h.Where("username = ?", username).First(o).Error
}

func (h *handler) UpdateUser(id uint, o *models.User) error {
	return h.Save(o).Error
	// return h.Model(models.User{}).Where("id=?", id).Update(o).Error
}

func (h *handler) DeleteUser(id uint) error {
	return h.Delete(models.User{ID: id}).Error
}

func (h *handler) RetrieveUserList(l *[]models.User, filter *UserFilter,
	orderBy *OrderBy, pag *Paginator) error {
	db := makeUserFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.User{}).TableName())
	return db.Find(l).Error
}

func (h *handler) CountUsers(filter *UserFilter) (count int, e error) {
	e = makeUserFilter(h.DB.Model(&models.User{}), filter).Count(&count).Error
	return
}

func makeUserFilter(db *gorm.DB, filter *UserFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.UsernamePrefix != nil {
		db = db.Where("username ILIKE ?", *(filter.UsernamePrefix)+"%")
	}
	if filter.Provider != nil {
		db = db.Where("provider = ?", *(filter.Provider))
	}
	if filter.Rol != nil {
		db = db.Where("rol=?", *(filter.Rol))
	}
	if filter.NameSubstr != nil {
		db = db.Where("name ILIKE ?", "%"+*(filter.NameSubstr)+"%")
	}
	if filter.Enabled != nil {
		db = db.Where("enabled = ?", *(filter.Enabled))
	}
	fmt.Println("AdminID filtering has to be doed by join to table group_admin")
	if filter.GroupWhichAdmin != nil {
		db = db.Where("id IN (?)",
			db.Model(&models.GroupAdmin{}).
				Select("user_id").Where("group_id = ?", *(filter.GroupWhichAdmin)).QueryExpr())
	}

	return db
}

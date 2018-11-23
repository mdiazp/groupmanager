package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// UserDBHandler ...
type UserDBHandler interface {
	CreateUser(o *models.User) error
	RetrieveUser(o *models.User) error
	RetrieveUserList(o *[]models.User, username *string, rol *models.Rol,
		adGroup *string) error
	UpdateUser(o *models.User) error
	DeleteUser(o *models.User) error
}

func (h *modelHandler) CreateUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveUserList(o *[]models.User, username *string, rol *models.Rol, adGroup *string) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) UpdateUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) DeleteUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// UserHandler ...
type UserHandler interface {
	CreateUser(o *models.User) error
	RetrieveUser(o *models.User) error
	RetrieveUserList(o *[]models.User, usernameSubstr *string, rol *models.Rol,
		adGroup *string, limit *int, offset *int, orderby *string, desc *bool) error
	UpdateUser(o *models.User) error
	DeleteUser(o *models.User) error

	CountUsers(usernameSubstr *string, rol *models.Rol, adGroup *string) (int, error)
}

func (h *handler) CreateUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveUserList(o *[]models.User, usernameSubstr *string,
	rol *models.Rol, adGroup *string, limit *int, offset *int, orderby *string, desc *bool) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) UpdateUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) DeleteUser(o *models.User) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) CountUsers(usernameSubstr *string, rol *models.Rol,
	adGroup *string) (int, error) {
	return 0, fmt.Errorf("not implemented yet")
}

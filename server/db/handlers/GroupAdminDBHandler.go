package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// GroupAdminHandler ...
type GroupAdminHandler interface {
	CreateGroupAdmin(o *models.GroupAdmin) error
	RetrieveGroupAdmin(o *models.GroupAdmin) error
	RetrieveGroupAdminList(o *[]models.GroupAdmin, groupID *int, userID *int,
		limit *int, offset *int, orderby *string, desc *bool) error
	UpdateGroupAdmin(o *models.GroupAdmin) error
	DeleteGroupAdmin(o *models.GroupAdmin) error

	CountGroupAdmin(groupID *int, userID *int) error
}

func (h *handler) CreateGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroupAdminList(o *[]models.GroupAdmin, groupID *int,
	userID *int, limit *int, offset *int, orderby *string, desc *bool) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) UpdateGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) DeleteGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) CountGroupAdmin(groupID *int, userID *int) error {
	return fmt.Errorf("not implemented yet")
}

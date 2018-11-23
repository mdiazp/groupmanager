package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// GroupAdminDBHandler ...
type GroupAdminDBHandler interface {
	CreateGroupAdmin(o *models.GroupAdmin) error
	RetrieveGroupAdmin(o *models.GroupAdmin) error
	RetrieveGroupAdminList(o *[]models.GroupAdmin, adGroup *string, userID *int) error
	UpdateGroupAdmin(o *models.GroupAdmin) error
	DeleteGroupAdmin(o *models.GroupAdmin) error
}

func (h *modelHandler) CreateGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveGroupAdminList(o *[]models.GroupAdmin, adGroup *string, userID *int) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) UpdateGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) DeleteGroupAdmin(o *models.GroupAdmin) error {
	return fmt.Errorf("not implemented yet")
}

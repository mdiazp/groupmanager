package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// GroupUserDBHandler ...
type GroupUserDBHandler interface {
	CreateGroupUser(o *models.GroupUser) error
	RetrieveGroupUser(o *models.GroupUser) error
	RetrieveGroupUserList(o *models.GroupUser, adGroup *string, adUser *string) error
	UpdateGroupUser(o *models.GroupUser) error
	DeleteGroupUser(o *models.GroupUser) error
}

func (h *modelHandler) CreateGroupUser(o *models.GroupUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveGroupUser(o *models.GroupUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) RetrieveGroupUserList(o *models.GroupUser, adGroup *string, adUser *string) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) UpdateGroupUser(o *models.GroupUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *modelHandler) DeleteGroupUser(o *models.GroupUser) error {
	return fmt.Errorf("not implemented yet")
}

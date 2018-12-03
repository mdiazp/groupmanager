package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// GroupADUserHandler ...
type GroupADUserHandler interface {
	CreateGroupADUser(o *models.GroupADUser) error
	RetrieveGroupADUser(o *models.GroupADUser) error
	RetrieveGroupADUserList(o *models.GroupADUser, groupID *int, adUserSubstr *string,
		adUserFullNameSubstr *string, limit *int, offset *int, orderby *string,
		desc *bool) error
	UpdateGroupADUser(o *models.GroupADUser) error
	DeleteGroupADUser(o *models.GroupADUser) error

	CountGroupADUser(groupID *int, adUserSubstr *string, adUserFullNameSubstr *string) (int, error)
}

func (h *handler) CreateGroupADUser(o *models.GroupADUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroupADUser(o *models.GroupADUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroupADUserList(o *models.GroupADUser, groupID *int,
	adUserSubstr *string, adUserFullNameSubstr *string, limit *int, offset *int, orderby *string,
	desc *bool) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) UpdateGroupADUser(o *models.GroupADUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) DeleteGroupADUser(o *models.GroupADUser) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) CountGroupADUser(groupID *int, adUserSubstr *string,
	adUserFullNameSubstr *string) (int, error) {
	return 0, fmt.Errorf("not implemented yet")
}

package handlers

import (
	"fmt"

	"github.com/mdiazp/groupmanager/server/db/models"
)

// GroupHandler ...
type GroupHandler interface {
	CreateGroup(o *models.Group) error
	RetrieveGroup(o *models.Group) error
	RetrieveGroupList(o *[]models.Group, GroupSubstr *string,
		adGroupSubstr *string, limit *int, offset *int, orderby *string, desc *bool) error
	UpdateGroup(o *models.Group) error
	DeleteGroup(o *models.Group) error

	CountGroups(GroupSubstr *string, adGroupSubstr *string) (int, error)
}

func (h *handler) CreateGroup(o *models.Group) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroup(o *models.Group) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) RetrieveGroupList(o *[]models.Group, GroupSubstr *string,
	adGroupSubstr *string, limit *int, offset *int, orderby *string, desc *bool) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) UpdateGroup(o *models.Group) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) DeleteGroup(o *models.Group) error {
	return fmt.Errorf("not implemented yet")
}

func (h *handler) CountGroups(GroupSubstr *string, adGroupSubstr *string) (int, error) {
	return 0, fmt.Errorf("not implemented yet")
}

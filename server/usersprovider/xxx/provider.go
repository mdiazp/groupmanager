package xxx

import (
	"errors"
	"fmt"

	"github.com/mdiazp/gm/server/usersprovider"
)

type provider struct{}

// Authenticate ...
func (p *provider) Authenticate(username, password string) error {
	if password != "123" {
		return errors.New("Fail Authentication")
	}
	return nil
}

func (p *provider) GetUserRecords(username string) (usersprovider.UserRecords, error) {
	return usersprovider.UserRecords{
		Username: username,
		Name:     username,
	}, nil
}

func (p *provider) GetFirst10BestUsernamePrefixMatchs(
	usernamePrefix string) (*[]usersprovider.UserRecords, error) {
	return nil, fmt.Errorf("Not Implemented")
}

// GetProvider ...
func GetProvider() usersprovider.Provider {
	return &provider{}
}

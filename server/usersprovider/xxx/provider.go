package xxx

import (
	"errors"

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

// GetProvider ...
func GetProvider() usersprovider.Provider {
	return &provider{}
}

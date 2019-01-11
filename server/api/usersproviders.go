package api

import (
	"fmt"

	"github.com/mdiazp/gm/server/usersprovider"
	"github.com/mdiazp/gm/server/usersprovider/ldap"
)

// UserProvider ...
type UserProvider string

const (
	// UserProviderAD ...
	UserProviderAD UserProvider = "AD"
	// UserProviderXXX ...
	UserProviderXXX UserProvider = "XXX"
	// UserProviderRoot ...
	UserProviderRoot UserProvider = "ROOT"
)

// GetUsersProviderNames ...
func (b *base) GetUsersProviderNames() []UserProvider {
	// return []UserProvider{UserProviderXXX, UserProviderAD}
	return []UserProvider{UserProviderAD, UserProviderRoot}
}

// GetUsersProvider ...
func (b *base) GetUsersProvider(provider UserProvider) usersprovider.Provider {
	switch provider {
	case UserProviderAD:
		return ldap.GetProvider(b.adConfig.AdAddress,
			b.adConfig.AdSuff,
			b.adConfig.AdBDN,
			b.adConfig.AdUser,
			b.adConfig.AdPassword)
	case UserProviderRoot:
		return &userRootProvider{
			b: b,
		}
	default:
		return nil
	}
}

// UserRootUsername ...
const UserRootUsername = "root"

type userRootProvider struct {
	b *base
}

func (p *userRootProvider) Authenticate(username, password string) (e error) {
	if username == UserRootUsername && p.b.userRootPassword == password {
		return nil
	}

	return fmt.Errorf("Fail authentication")
}

func (p *userRootProvider) GetUserRecords(username string) (usersprovider.UserRecords, error) {
	if username == UserRootUsername {
		return usersprovider.UserRecords{
			Username: UserRootUsername,
			Name:     "Root",
		}, nil
	}
	return usersprovider.UserRecords{}, fmt.Errorf("User Not Found")
}

func (p *userRootProvider) GetFirst10BestUsernamePrefixMatchs(
	usernamePrefix string) (*[]usersprovider.UserRecords, error) {
	return nil, fmt.Errorf("Not Implemented")
}

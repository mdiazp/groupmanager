package api

import (
	"github.com/mdiazp/gm/server/usersprovider"
	"github.com/mdiazp/gm/server/usersprovider/xxx"
)

// UserProvider ...
type UserProvider string

const (
	// UserProviderAD ...
	UserProviderAD UserProvider = "AD"
	// UserProviderXXX ...
	UserProviderXXX UserProvider = "XXX"
)

// GetUsersProviderNames ...
func GetUsersProviderNames() []UserProvider {
	return []UserProvider{UserProviderXXX, UserProviderAD}
}

// GetUsersProvider ...
func GetUsersProvider(b Base, provider UserProvider) usersprovider.Provider {
	switch provider {
	case UserProviderXXX:
		if b.GetEnv() == "dev" {
			return xxx.GetProvider()
		}
		return nil
	case UserProviderAD:
		if b.GetEnv() == "dev" {
			return xxx.GetProvider()
		}
		return nil
	default:
		return nil
	}
}

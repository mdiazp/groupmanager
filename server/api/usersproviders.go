package api

import (
	"github.com/mdiazp/gm/server/usersprovider"
	"github.com/mdiazp/gm/server/usersprovider/ldap"
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
func (b *base) GetUsersProviderNames() []UserProvider {
	// return []UserProvider{UserProviderXXX, UserProviderAD}
	return []UserProvider{UserProviderAD}
}

// GetUsersProvider ...
func (b *base) GetUsersProvider(provider UserProvider) usersprovider.Provider {
	switch provider {
	case UserProviderXXX:
		if b.GetEnv() == "dev" {
			return xxx.GetProvider()
		}
		return nil
	case UserProviderAD:
		if b.GetEnv() == "dev" {
			return ldap.GetProvider(b.adConfig.AdAddress,
				b.adConfig.AdSuff,
				b.adConfig.AdBDN,
				b.adConfig.AdUser,
				b.adConfig.AdPassword)
		}
		return ldap.GetProvider(b.adConfig.AdAddress,
			b.adConfig.AdSuff,
			b.adConfig.AdBDN,
			b.adConfig.AdUser,
			b.adConfig.AdPassword)
	default:
		return nil
	}
}

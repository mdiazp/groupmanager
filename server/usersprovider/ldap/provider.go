package ldap

import (
	"github.com/lamg/ldaputil"
	"github.com/mdiazp/gm/server/usersprovider"
)

type provider struct {
	ldap *ldaputil.Ldap
}

// Authenticate ...
func (p *provider) Authenticate(username, password string) error {
	e := p.ldap.Authenticate(username, password)
	return e
}

// GetUserRecords ...
func (p *provider) GetUserRecords(username string) (usersprovider.UserRecords, error) {
	m, e := p.ldap.FullRecordAcc(username)

	if e != nil {
		return zero, e
	}

	u := zero
	setv := func(x []string, v *string) {
		*v = ""
		if x != nil && len(x) > 0 {
			*v = x[0]
		}
	}
	setv(m["sAMAccountName"], &u.Username)
	setv(m["displayName"], &u.Name)

	return u, e
}

var zero usersprovider.UserRecords

// GetProvider ...
func GetProvider(AdAddress, AdSuff, AdBDN, AdUser, AdPassword string) usersprovider.Provider {
	provider := &provider{
		ldap: ldaputil.NewLdapWithAcc(
			AdAddress,
			AdSuff,
			AdBDN,
			AdUser,
			AdPassword,
		),
	}

	return provider
}

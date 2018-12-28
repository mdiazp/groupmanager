package ldap

import (
	"github.com/mdiazp/sirel-server/api/pkg/authproviders"
	"github.com/mdiazp/sirel-server/api/pkg/ldaputil"
)

type provider struct {
	ldap *ldaputil.Ldap
}

func (this *provider) Authenticate(username, password string) error {
	e := this.ldap.Authenticate(username, password)
	return e
}

func (this *provider) GetUserRecords(username string) (authproviders.UserRecords, error) {
	m, e := this.ldap.FullRecordAcc(username)

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
	setv(m["mail"], &u.Email)

	return u, e
}

var zero authproviders.UserRecords

func GetProvider(AdAddress, AdSuff, AdBDN, AdUser, AdPassword string) authproviders.Provider {
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

package ldap

// Handler ...
type Handler interface {
}

// ADConfig ...
type ADConfig interface {
	GetAdAddress() string
	GetAdSuff() string
	GetAdBDN() string
	GetAdUser() string
	GetAdPassword() string
}

// NewHandler ...
func NewHandler(config ADConfig) (Handler, error) {
	return new(handler), nil
}

type handler struct {
}

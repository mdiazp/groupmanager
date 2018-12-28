package usersprovider

// UserRecords ...
type UserRecords struct {
	Username string
	Name     string
}

// Provider ...
type Provider interface {
	Authenticate(username, password string) (e error)
	GetUserRecords(username string) (UserRecords, error)
}

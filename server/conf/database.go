package conf

// DatabaseConfig ...
type DatabaseConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBDriver   string
}

// GetDBName ...
func (d *DatabaseConfig) GetDBName() string {
	return d.DBName
}

// GetDBUser ...
func (d *DatabaseConfig) GetDBUser() string {
	return d.DBUser
}

// GetDBPassword ...
func (d *DatabaseConfig) GetDBPassword() string {
	return d.DBPassword
}

// GetDBDriver ...
func (d *DatabaseConfig) GetDBDriver() string {
	return d.DBDriver
}

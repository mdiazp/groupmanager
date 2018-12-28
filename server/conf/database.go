package conf

// DatabaseConfig ...
type DatabaseConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBDialect  string
}

// GetDBHost ...
func (d *DatabaseConfig) GetDBHost() string {
	return d.DBHost
}

// GetDBPort ...
func (d *DatabaseConfig) GetDBPort() int {
	return d.DBPort
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

// GetDBDialect ...
func (d *DatabaseConfig) GetDBDialect() string {
	return d.DBDialect
}

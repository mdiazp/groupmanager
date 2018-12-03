package conf

// ADConfig ...
type ADConfig struct {
	AdAddress  string
	AdSuff     string
	AdBDN      string
	AdUser     string
	AdPassword string
}

// GetAdAddress ...
func (c *ADConfig) GetAdAddress() string {
	return c.AdAddress
}

// GetAdSuff ...
func (c *ADConfig) GetAdSuff() string {
	return c.AdSuff
}

// GetAdBDN ...
func (c *ADConfig) GetAdBDN() string {
	return c.AdBDN
}

// GetAdUser ...
func (c *ADConfig) GetAdUser() string {
	return c.AdUser
}

// GetAdPassword ...
func (c *ADConfig) GetAdPassword() string {
	return c.AdPassword
}

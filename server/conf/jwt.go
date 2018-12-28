package conf

// JWTConfig ...
type JWTConfig struct {
	Secret string
}

// GetSecret ...
func (c *JWTConfig) GetSecret() string {
	return c.Secret
}

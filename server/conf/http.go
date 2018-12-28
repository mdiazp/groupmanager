package conf

// HTTPConfig ...
type HTTPConfig struct {
	Host           string
	Port           string
	ReadTimeout    int64
	WriteTimeout   int64
	MaxHeaderBytes int
}

package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration ...
type Configuration struct {
	DatabaseConfig
	LogsConfig
	JWTConfig
	ADConfig
	HTTPConfig
	APIDoc

	UserRootConfig
}

// LoadConfiguration ...
func LoadConfiguration(path string, env string) (*Configuration, error) {
	file, err := os.Open(path + "/config." + env + ".json")
	if err != nil {
		return nil, fmt.Errorf("Fail openning configuration file: %s", err.Error())
	}
	defer file.Close()

	config := &Configuration{}

	//Parsing json file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("Configuration file has wrong format: %s", err.Error())
	}

	return config, nil
}

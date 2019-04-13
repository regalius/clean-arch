package config

import (
	"errors"
	"log"
)

// Get get configuration value from an instance
func Get(configName, key string) (value interface{}, err error) {
	if vInstances[configName] == nil {
		log.Println("[Common/Config] Viper instance isn't initialized for", configName)
		err = errors.New("Viper instance isn't initialized")
		return
	}

	value = vInstances[configName].Get(key)
	return
}

package config

import (
	"bytes"
	"errors"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// NewConfigFromFile instantiate new configuration instance
// Config in this example is read-only, you can't set it on the fly
// so don't return any possible interface to do so
func NewConfigFromFile(configName, configType, filePath string, options NewConfigOptions) (err error) {
	if vInstances[configName] != nil {
		log.Println("[Common/Config] Viper instance is already initialized for", configName)
		err = errors.New("Viper instance is already initialized")
		return
	}
	vInstances[configName] = viper.New()
	vInstances[configName].SetConfigType(configType)
	vInstances[configName].SetConfigFile(filePath)

	if options.Defaults != nil {
		defaults := options.Defaults
		for key, value := range defaults {
			vInstances[configName].SetDefault(key, value)
		}
	}

	err = vInstances[configName].ReadInConfig()
	if err != nil {
		log.Println("[Common/Config] Can't read config file,", err)
		return
	}

	if options.IsWatch {
		vInstances[configName].WatchConfig()
		vInstances[configName].OnConfigChange(func(e fsnotify.Event) {
			log.Println("[Common/Config] Config File Changed", e.Name)
		})
	}

	return
}

// NewConfigFromString instantiate new configuration instance using string, useful for testing
func NewConfigFromString(configName, value string, options NewConfigOptions) (err error) {
	if vInstances[configName] != nil {
		log.Println("[Common/Config] Viper instance is already initialized for", configName)
		err = errors.New("Viper instance is already initialized")
		return
	}
	vInstances[configName] = viper.New()
	valueByte := []byte(value)
	viper.ReadConfig(bytes.NewBuffer(valueByte))

	if options.Defaults != nil {
		defaults := options.Defaults
		for key, value := range defaults {
			vInstances[configName].SetDefault(key, value)
		}
	}
	return
}

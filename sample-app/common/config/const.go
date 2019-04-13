package config

import (
	"github.com/spf13/viper"
)

var vInstances = make(map[string]*viper.Viper)

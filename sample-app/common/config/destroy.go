package config

import "log"

// Destroy deletes current configuration instance (useful for testing)
func Destroy(name string) {
	vInstances[name] = nil
	log.Println("[Common/Config] Config instance destroyed for", name)
	return
}

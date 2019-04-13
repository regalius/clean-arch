package main

var (
	fileList = map[string]string{
		"product":  "./configs/var/data/product.json",
		"user":     "./configs/var/data/user.json",
		"userPA":   "./configs/var/data/user-product-affinity.json",
		"genderPA": "./configs/var/data/gender-product-affinity.json",
	}
	configPaths = map[string]string{
		"production":  "/etc/recom-engine/production.toml",
		"development": "configs/etc/recom-engine/development.toml",
	}
)

const (
	configName string = "recomengine"
)

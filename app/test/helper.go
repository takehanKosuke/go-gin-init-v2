package test

import "go-gin-init-v2/app/config"

// LoadConfig
func LoadConfig() *config.Config {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}
	return config
}

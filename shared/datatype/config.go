package datatype

import (
	"os"
)

type Config struct {
	UserServiceURL     string
	CategoryServiceURL string
}

var config *Config

func NewConfig() *Config {
	if config == nil {
		config = &Config{
			UserServiceURL:     os.Getenv("USER_SERVICE_URL"),
			CategoryServiceURL: os.Getenv("CATEGORY_SERVICE_URL"),
		}
	}
	
	return config
}

func GetConfig() *Config {
	return NewConfig()
}

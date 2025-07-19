package datatype

import "os"

type AWSConfig struct {
	AccessKey  string
	BucketName string
	Domain     string
	Region     string
	SecretKey  string
}

type Config struct {
	UserServiceURL     string
	CategoryServiceURL string
	AWS                AWSConfig
}

var config *Config

func NewConfig() *Config {
	if config == nil {
		config = &Config{
			UserServiceURL:     os.Getenv("USER_SERVICE_URL"),
			CategoryServiceURL: os.Getenv("CATEGORY_SERVICE_URL"),
			AWS: AWSConfig{
				AccessKey:  os.Getenv("AWS_ACCESS_KEY"),
				BucketName: os.Getenv("AWS_BUCKET_NAME"),
				Domain:     os.Getenv("AWS_DOMAIN"),
				Region:     os.Getenv("AWS_REGION"),
				SecretKey:  os.Getenv("AWS_SECRET_KEY"),
			},
		}
	}
	
	return config
}

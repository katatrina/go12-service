package datatype

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AWSConfig struct {
	AccessKey  string `mapstructure:"AWS_ACCESS_KEY"`
	BucketName string `mapstructure:"AWS_BUCKET_NAME"`
	Domain     string `mapstructure:"AWS_DOMAIN"`
	Region     string `mapstructure:"AWS_REGION"`
	SecretKey  string `mapstructure:"AWS_SECRET_KEY"`
}

type GrpcConfig struct {
	CategoryServiceURL   string `mapstructure:"CATEGORY_GRPC_URL"`
	FoodServiceURL       string `mapstructure:"FOOD_GRPC_URL"`
	UserServiceURL       string `mapstructure:"USER_GRPC_URL"`
	RestaurantServiceURL string `mapstructure:"RESTAURANT_GRPC_URL"`
}

type Config struct {
	Port         string      `mapstructure:"PORT"`
	DBDSN        string      `mapstructure:"DB_DSN"`
	JWTSecretKey string      `mapstructure:"JWT_SECRET_KEY"`
	NatsURL      string      `mapstructure:"NATS_URL"`
	AWS          AWSConfig   `mapstructure:",squash"`
	Grpc         GrpcConfig  `mapstructure:",squash"`
}

var config *Config

func NewConfig() *Config {
	if config == nil {
		// Initialize viper
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./")

		// Enable reading from environment variables
		viper.AutomaticEnv()

		// Set default values
		setDefaults()

		// Try to read config file
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Could not read .env file: %v. Using environment variables and defaults.", err)
		}

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			log.Fatalf("Unable to decode config: %v", err)
		}
	}
	
	return config
}

func setDefaults() {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("CATEGORY_GRPC_URL", "localhost:6000")
	viper.SetDefault("FOOD_GRPC_URL", "localhost:6001")
	viper.SetDefault("USER_GRPC_URL", "localhost:6002")
	viper.SetDefault("RESTAURANT_GRPC_URL", "localhost:6003")
	viper.SetDefault("NATS_URL", "nats://127.0.0.1:4222")
	viper.SetDefault("AWS_REGION", "ap-southeast-1")
}


func (config *Config) GetConfig() *Config {
	return config
}

// GetPortFromURL extracts port from URL string (e.g., "localhost:6000" -> "6000")
func (gc *GrpcConfig) GetCategoryPort() string {
	return extractPort(gc.CategoryServiceURL)
}

func (gc *GrpcConfig) GetFoodPort() string {
	return extractPort(gc.FoodServiceURL)
}

func (gc *GrpcConfig) GetUserPort() string {
	return extractPort(gc.UserServiceURL)
}

func (gc *GrpcConfig) GetRestaurantPort() string {
	return extractPort(gc.RestaurantServiceURL)
}

func extractPort(url string) string {
	parts := strings.Split(url, ":")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return "8080" // fallback
}

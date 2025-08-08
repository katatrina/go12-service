package datatype

import (
	"log"
	"os"

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
	Port                  string `mapstructure:"GRPC_PORT"`
	FoodServicePort       string `mapstructure:"FOOD_SERVICE_GRPC_PORT"`
	UserServicePort       string `mapstructure:"USER_SERVICE_GRPC_PORT"`
	RestaurantServicePort string `mapstructure:"RESTAURANT_SERVICE_GRPC_PORT"`
	CategoryServiceURL    string `mapstructure:"CATEGORY_SERVICE_GRPC_URL"`
	FoodServiceURL        string `mapstructure:"FOOD_SERVICE_GRPC_URL"`
	UserServiceURL        string `mapstructure:"USER_SERVICE_GRPC_URL"`
	RestaurantServiceURL  string `mapstructure:"RESTAURANT_SERVICE_GRPC_URL"`
}

type Config struct {
	Port               string      `mapstructure:"PORT"`
	DBDSN              string      `mapstructure:"DB_DSN"`
	UserServiceURL     string      `mapstructure:"USER_SERVICE_URL"`
	CategoryServiceURL string      `mapstructure:"CATEGORY_SERVICE_URL"`
	FoodServiceURL     string      `mapstructure:"FOOD_SERVICE_URL"`
	JWTSecretKey       string      `mapstructure:"JWT_SECRET_KEY"`
	NatsURL            string      `mapstructure:"NATS_URL"`
	AWS                AWSConfig   `mapstructure:",squash"`
	Grpc               GrpcConfig  `mapstructure:",squash"`
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

		// Fallback to os.Getenv for backward compatibility
		fallbackToEnvVars(config)
	}
	
	return config
}

func setDefaults() {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("GRPC_PORT", "6000")
	viper.SetDefault("FOOD_SERVICE_GRPC_PORT", "6001")
	viper.SetDefault("USER_SERVICE_GRPC_PORT", "6002")
	viper.SetDefault("RESTAURANT_SERVICE_GRPC_PORT", "6003")
	viper.SetDefault("CATEGORY_SERVICE_GRPC_URL", "localhost:6000")
	viper.SetDefault("FOOD_SERVICE_GRPC_URL", "localhost:6001")
	viper.SetDefault("USER_SERVICE_GRPC_URL", "localhost:6002")
	viper.SetDefault("RESTAURANT_SERVICE_GRPC_URL", "localhost:6003")
	viper.SetDefault("NATS_URL", "nats://127.0.0.1:4222")
	viper.SetDefault("AWS_REGION", "ap-southeast-1")
}

func fallbackToEnvVars(cfg *Config) {
	if cfg.DBDSN == "" {
		cfg.DBDSN = os.Getenv("DB_DSN")
	}
	if cfg.Port == "" {
		cfg.Port = os.Getenv("PORT")
		if cfg.Port == "" {
			cfg.Port = "8080"
		}
	}
	if cfg.UserServiceURL == "" {
		cfg.UserServiceURL = os.Getenv("USER_SERVICE_URL")
	}
	if cfg.CategoryServiceURL == "" {
		cfg.CategoryServiceURL = os.Getenv("CATEGORY_SERVICE_URL")
	}
	if cfg.FoodServiceURL == "" {
		cfg.FoodServiceURL = os.Getenv("FOOD_SERVICE_URL")
	}
	if cfg.JWTSecretKey == "" {
		cfg.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	}
	if cfg.NatsURL == "" {
		cfg.NatsURL = os.Getenv("NATS_URL")
	}
	
	// AWS fallbacks
	if cfg.AWS.AccessKey == "" {
		cfg.AWS.AccessKey = os.Getenv("AWS_ACCESS_KEY")
	}
	if cfg.AWS.BucketName == "" {
		cfg.AWS.BucketName = os.Getenv("AWS_BUCKET_NAME")
	}
	if cfg.AWS.Domain == "" {
		cfg.AWS.Domain = os.Getenv("AWS_DOMAIN")
	}
	if cfg.AWS.Region == "" {
		cfg.AWS.Region = os.Getenv("AWS_REGION")
	}
	if cfg.AWS.SecretKey == "" {
		cfg.AWS.SecretKey = os.Getenv("AWS_SECRET_KEY")
	}
	
	// gRPC fallbacks
	if cfg.Grpc.Port == "" {
		cfg.Grpc.Port = os.Getenv("GRPC_PORT")
		if cfg.Grpc.Port == "" {
			cfg.Grpc.Port = "6000"
		}
	}
	if cfg.Grpc.FoodServicePort == "" {
		cfg.Grpc.FoodServicePort = os.Getenv("FOOD_SERVICE_GRPC_PORT")
		if cfg.Grpc.FoodServicePort == "" {
			cfg.Grpc.FoodServicePort = "6001"
		}
	}
	if cfg.Grpc.CategoryServiceURL == "" {
		cfg.Grpc.CategoryServiceURL = os.Getenv("CATEGORY_SERVICE_GRPC_URL")
		if cfg.Grpc.CategoryServiceURL == "" {
			cfg.Grpc.CategoryServiceURL = "localhost:6000"
		}
	}
	if cfg.Grpc.FoodServiceURL == "" {
		cfg.Grpc.FoodServiceURL = os.Getenv("FOOD_SERVICE_GRPC_URL")
		if cfg.Grpc.FoodServiceURL == "" {
			cfg.Grpc.FoodServiceURL = "localhost:6001"
		}
	}
}

func (config *Config) GetConfig() *Config {
	return config
}

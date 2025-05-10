package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	AppConfig = &Config{
		Port:   viper.GetString("PORT"),
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASSWORD"),
		DBName: viper.GetString("DB_NAME"),
	}
}

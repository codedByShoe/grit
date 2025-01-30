package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig struct {
		Url      string `mapstructure:"DB_URL"`
		Port     int    `mapstructure:"DB_PORT"`
		Host     string `mapstructure:"DB_HOST"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		SSLMode  string `mapstructure:"DB_SSLMODE"`
	}

	SessionConfig struct {
		SecretKey string `mapstructure:"SECRET_KEY"`
	}

	ServerConfig struct {
		Url  string `mapstructure:"URL"`
		Port int    `mapstructure:"PORT"`
		Env  string `mapstructure:"ENVIRONMENT"`
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("PORT", 8000)
	viper.SetDefault("URL", "http://localhost")
	viper.SetDefault("DB_URL", "sqlite3:///database.sqlite")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

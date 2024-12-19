package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "codedbyshoe",
		DBPassword: "password",
		DBName:     "gritdb",
		ServerPort: "8080",
	}
}

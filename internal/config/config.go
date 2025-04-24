package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP struct {
		Port         string        `env:"HTTP_PORT" env-default:"8080"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" env-default:"10s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" env-default:"10s"`
		BasePath     string        `env:"HTTP_BASE_PATH" env-default:""`
	}
	MongoDB struct {
		URI      string        `env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
		Database string        `env:"MONGO_DB" env-default:"tusurprk"`
		Timeout  time.Duration `env:"MONGO_TIMEOUT" env-default:"5s"`
	}
	Auth struct {
		JWTSecret   string `env:"JWT_SECRET" env-required:"true"`
		UserIDClaim string `env:"USER_ID_CLAIM" env-default:"user_id"`
	}
	Storage struct {
		UploadDir   string `env:"UPLOAD_DIR" env-default:"./uploads"`
		MaxFileSize int64  `env:"MAX_FILE_SIZE" env-default:"10485760"`
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env:", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфига: %v", err) // Детализируем ошибку
	}

	return cfg, nil
}

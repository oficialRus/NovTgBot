package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Telegram struct {
		Token string `env:"TOKEN_TG" env-required:"true"`
	}
}

func MustLoad() Config {
	// Загружаем переменные из .env файла
	// Игнорируем ошибку, если файл не найден (переменные могут быть установлены в системе)
	_ = godotenv.Load()

	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	return cfg
}

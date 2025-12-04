package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Telegram struct {
		Token string `env:"TOKEN_TG" env-required:"true"`
	}
}

func MustLoad() Config {
	var cfg Config
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Info
	}

}

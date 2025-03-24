package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Conf struct {
	Port     string `env:"MYSQL_PORT" env-default:"3306"`
	Host     string `env:"MYSQL_HOST" env-default:"localhost"`
	Name     string `env:"MYSQL_DATABASE"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}

func MustLoad() (*Conf, error) {
	var Cfg Conf

	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		return nil, fmt.Errorf("error parse required params: %s", err)
	}

	return &Cfg, nil
}

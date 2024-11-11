package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mkorobovv/full-restaurant/internal/app/config"
)

func New() (config config.Config) {
	err := cleanenv.ReadConfig("./config.yml", &config)
	if err != nil {
		err = fmt.Errorf(strings.ReplaceAll(err.Error(), ", ", ",\n"))

		log.Fatalln(err)
	}

	return
}

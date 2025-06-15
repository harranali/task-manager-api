package config

import (
	"log"
	"strconv"

	"github.com/harranali/task-manager-api/utils"
)

type Config struct {
	JWTSecret   string
	JWTDuration int
}

var CFG *Config

func NewConfig() *Config {
	jwtDuration, err := strconv.Atoi(utils.GetEnvMust("JWT_DURATION_HOURS"))
	if err != nil {
		log.Fatal("unable to parse to int")
	}
	CFG = &Config{
		JWTSecret:   utils.GetEnvMust("JWT_SECRET"),
		JWTDuration: jwtDuration,
	}
	return CFG
}

package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTP     serverConfig `envconfig:"HTTP"`
	Postgres dbConfig     `envconfig:"POSTGRES"`
	JWT      jwtConfig    `envconfig:"JWT"`
}

func configFromEnv() (*config, error) {
	var c config
	err := envconfig.Process("CHATBOOK", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type dbConfig struct {
	Host         string `envconfig:"HOST" default:"localhost"`
	Port         int    `envconfig:"PORT" default:"5433"`
	User         string `envconfig:"USER" default:"um-dev"`
	Password     string `envconfig:"PASSWORD" default:"123456"`
	DBName       string `envconfig:"DB_NAME" default:"um-dev-db"`
	MaxIdleConns int    `envconfig:"MAX_IDLE_CONNS" default:"16"`
	MaxOpenConns int    `envconfig:"MAX_OPEN_CONNS" default:"32"`
	Mode         string `envconfig:"MODE" default:"DEV"`
}

type serverConfig struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"` //localhost
	Port int    `envconfig:"PORT" default:"8080"`
}

type jwtConfig struct {
	Issuer string `envconfig:"ISSUER"`
	Secret string `envconfig:"JWT_SECRET"`
}

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
	Host         string `envconfig:"HOST"`
	Port         int    `envconfig:"PORT"`
	User         string `envconfig:"USER"`
	Password     string `envconfig:"PASSWORD"`
	DBName       string `envconfig:"DB_NAME"`
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

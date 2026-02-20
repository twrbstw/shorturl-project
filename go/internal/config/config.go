package config

import "os"

type Configs struct {
	DbConfig DbConfig
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDefaultConfig() Configs {
	return Configs{
		DbConfig: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}

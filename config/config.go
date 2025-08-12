package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	TIMEZONE    string
}

func Config() Conf {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	return Conf{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		TIMEZONE:    os.Getenv("TIMEZONE"),
	}
}

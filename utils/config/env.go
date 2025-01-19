package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Mode                     string
	DB_TYPE                  string
	DB_USER                  string
	DB_PASSWORD              string
	DB_HOST                  string
	DB_PORT                  string
	DB_NAME                  string
	SSLMODE                  string
	INSTANCE_CONNECTION_NAME string

	DB_DSN string
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load the .env: %v", err)
		return err
	}

	Mode = os.Getenv("MODE")
	DB_TYPE = os.Getenv("DB_TYPE")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	SSLMODE = os.Getenv("SSLMODE")
	INSTANCE_CONNECTION_NAME = os.Getenv("INSTANCE_CONNECTION_NAME")

	DB_DSN = "user=" + DB_USER +
		" password=" + DB_PASSWORD +
		" host=" + DB_HOST +
		" port=" + DB_PORT +
		" dbname=" + DB_NAME +
		" sslmode=" + SSLMODE

	return nil
}

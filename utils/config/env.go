package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbUser                 string
	dbPassword             string
	dbName                 string
	instabceConnectionName string
	privateIP              string

	Mode   string
	DBType string
	DBDSN  string
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load the .env: %v", err)
		return err
	}

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	instabceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	privateIP = os.Getenv("PRIVATE_IP")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSLMODE")

	Mode = os.Getenv("MODE")
	DBType = os.Getenv("DB_TYPE")
	DBDSN = "user=" + dbUser +
		" password=" + dbPassword +
		" host=" + dbHost +
		" port=" + dbPort +
		" dbname=" + dbName +
		" sslmode=" + sslmode

	return nil
}

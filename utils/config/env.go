package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/polupolu-dev/polupolu-backend/utils/consts"
)

var (
	DBName                 string
	InstanceConnectionName string
	PrivateIP              string

	Mode   string
	DBType string
	DBDSN  string
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load the .env: %v", err)
		return err
	}

	DBName = os.Getenv("DB_NAME")
	InstanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	PrivateIP = os.Getenv("PRIVATE_IP")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSLMODE")

	Mode = os.Getenv("MODE")
	DBType = os.Getenv("DB_TYPE")

	switch DBType {
	case consts.Cloudsql:
		DBDSN = "user=" + dbUser +
			" password=" + dbPassword +
			" database=" + DBName
	case consts.Postgres:
		DBDSN = "user=" + dbUser +
			" password=" + dbPassword +
			" host=" + dbHost +
			" port=" + dbPort +
			" dbname=" + DBName +
			" sslmode=" + sslmode
	}

	return nil
}

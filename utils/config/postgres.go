package config

import (
	"database/sql"
	"log"
)

func Postgres() *sql.DB {
	dbConn, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return dbConn
}

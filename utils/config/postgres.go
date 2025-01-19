package config

import (
	"database/sql"
	"fmt"
)

func Postgres() (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", DBDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return dbConn, err
}

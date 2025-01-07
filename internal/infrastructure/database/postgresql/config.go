package postgresql

import (
	"database/sql"
)

type PostgresqlDependency struct {
	db *sql.DB
}

func (d *PostgresqlDependency) OpenDB() (*sql.DB, error) {
	// PostgreSQL データベースファイルのオープン
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return db, err
	}
	return db, nil
}

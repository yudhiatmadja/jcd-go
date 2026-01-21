package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DB_DSN"))
}

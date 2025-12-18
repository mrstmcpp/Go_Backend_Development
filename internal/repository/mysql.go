package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
	dbUrl := "postgres://admin:admin@localhost:5432/go_backend_dev?sslmode=disable"
	return sql.Open("postgres", dbUrl)
}

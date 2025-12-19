package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {

	dbUrl := os.Getenv("DATABASE_URL")

	// fmt.Println("DATABASE_URL =", os.Getenv("DATABASE_URL"))

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

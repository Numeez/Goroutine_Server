package main

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl==""{
		return nil,errors.New("DB_URL env variable is missing")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err

	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

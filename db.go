package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func connectDB() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("DB_URL env variable is missing")
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

func (s Server) makeUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		surname VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := s.dbConnection.Exec(query)
	if err != nil {
		return err
	}
	return nil

}

func (s Server) insertUser(user User) error {
	query := `INSERT INTO users (name, surname) VALUES ($1, $2)`
	_, err := s.dbConnection.Exec(query, user.Name, user.Surname)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (s Server) getAllUsers() ([]User, error) {
	result := []User{}
	query := `SELECT name,surname FROM users`
	rows, err := s.dbConnection.Query(query)
	if err != nil {
		return result, fmt.Errorf("failed to insert user: %w", err)
	}
	for rows.Next() {
		var name, surname string
		err := rows.Scan(&name, &surname)
		if err != nil {
			return result, err
		}
		user := User{
			Name:    name,
			Surname: surname,
		}
		result = append(result, user)

	}
	return result, nil
}

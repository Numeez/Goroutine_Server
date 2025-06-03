package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer(":8080", db)
	server.run()
}

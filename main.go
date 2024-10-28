package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	DB_FILE := os.Getenv("DB_FILE")
	PORT := ":" + os.Getenv("PORT")

	db, err := NewSQLiteStorage(DB_FILE)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(PORT, db)
	server.Run()
}

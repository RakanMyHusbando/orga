package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in .env file")
	}

	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		log.Fatal("DB_FILE must be set in .env file")
	}

	server := NewAPIServer(":" + port)
	server.Run()

	db, err := NewSQLiteStorage(dbFile)
	CreateSQLiteTable(db)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
)

func main() {
	const PORT, DB_FILE string = "8080", "data.db"

	// server := NewAPIServer(":" + PORT)
	// server.Run()

	_, err := NewSQLiteStorage(DB_FILE)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import "log"

func main() {
	const PORT, DB_FILE string = "8080", "data.db"

	db, err := NewSQLiteStorage(DB_FILE)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":"+PORT, db)
	server.Run()
}

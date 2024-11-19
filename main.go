package main

import (
	"log"
	"os"

	"github.com/RakanMyHusbando/shogun/api"
	"github.com/RakanMyHusbando/shogun/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	DB_FILE := os.Getenv("DB_FILE")
	PORT := ":" + os.Getenv("PORT")

	db, err := storage.NewSQLiteStorage(DB_FILE)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.RunSQLiteStorage(db, "schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(PORT, db)
	server.Run()
}

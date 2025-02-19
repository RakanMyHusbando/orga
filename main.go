package main

import (
	"log"
	"os"

	"github.com/RakanMyHusbando/orga/api"
	"github.com/RakanMyHusbando/orga/storage"
	"github.com/joho/godotenv"
)

var (
	apiKey  string
	dbFile  string
	baseUrl string
)

func main() {
	godotenv.Load(".env")
	apiKey = os.Getenv("API_KEY")
	dbFile = os.Getenv("DB_FILE")
	baseUrl = os.Getenv("API_PORT")
	if baseUrl == "" {
		log.Fatal("Missing API_PORT in .env file")
	}
	baseUrl = os.Getenv("API_HOST") + ":" + baseUrl
	if os.Getenv("API_HOST") == "" {
		baseUrl = "127.0.0.1" + baseUrl
	}

	db, err := storage.NewSQLiteStorage(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.RunSQLiteStorage(db, "schema.sql", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(baseUrl, db)
	server.Run()
}

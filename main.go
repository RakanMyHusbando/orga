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
	env := map[string]string{
		"API_KEY": "",
		"DB_FILE": "",
		"PORT":    ":",
	}
	for key := range env {
		env[key] += os.Getenv(key)
		if env[key] == "" || (key == "PORT" && env[key] == ":") {
			log.Fatalf("Error getting %v", key)
		}
	}

	db, err := storage.NewSQLiteStorage(env["DB_FILE"])
	if err != nil {
		log.Fatal(err)
	}

	err = storage.RunSQLiteStorage(db, "schema.sql", env["API_KEY"])
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(env["PORT"], db)
	server.Run()
}

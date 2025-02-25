package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RakanMyHusbando/orga/api"
	"github.com/RakanMyHusbando/orga/storage"
	"github.com/RakanMyHusbando/orga/website"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	apiKey         string
	baseUrl        string
	domain         string
	dcClientSecret string
	dcClientId     string
	router         = mux.NewRouter()
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed reading environment variables: %v", err)
	}

	if err := runStorageAndApi(); err != nil {
		log.Fatalf("Failed to run storage and api: %v", err)
	}

	if err := runWebsite(); err != nil {
		log.Fatalf("Failed to run website: %v", err)
	}

	log.Println("Server running on", domain)
	log.Fatal(http.ListenAndServe(baseUrl, router))
}

func loadEnv() error {
	godotenv.Load(".env")
	domain = "http://" + os.Getenv("DOMAIN")
	apiKey = os.Getenv("API_KEY")
	baseUrl = os.Getenv("PORT")
	dcClientSecret = os.Getenv("DC_CLIENT_SECRET")
	dcClientId = os.Getenv("DC_CLIENT_ID")
	if baseUrl == "" {
		return fmt.Errorf("Missing PORT in .env file")
	}
	baseUrl = os.Getenv("HOST") + ":" + baseUrl
	if os.Getenv("HOST") == "" {
		baseUrl = "127.0.0.1" + baseUrl
	}
	if domain == "http://" {
		domain = "http://" + baseUrl
	}
	return nil
}

func runStorageAndApi() error {
	db, err := storage.NewSQLiteStorage()
	if err != nil {
		return err
	}
	if err = storage.RunSQLiteStorage(db, "schema.sql", apiKey); err != nil {
		return err
	}
	api.NewStore(db).Routes(router.PathPrefix("/api").Subrouter())
	return nil
}

func runWebsite() error {
	storage, err := website.NewSessionStorage()
	if err != nil {
		return err
	}
	web, err := website.NewWebsite(storage, dcClientId, dcClientSecret, domain)
	if err != nil {
		return err
	}
	web.Routes(router.PathPrefix("/").Subrouter())
	return nil
}

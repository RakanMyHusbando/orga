package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RakanMyHusbando/orga/httpHandler"
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
)

func main() {
	godotenv.Load(".env")
	domain = os.Getenv("DOMAIN")
	apiKey = os.Getenv("API_KEY")
	baseUrl = os.Getenv("PORT")
	dcClientSecret = os.Getenv("DC_CLIENT_SECRET")
	dcClientId = os.Getenv("DC_CLIENT_ID")
	if baseUrl == "" {
		log.Fatal("Missing PORT in .env file")
	}
	baseUrl = os.Getenv("HOST") + ":" + baseUrl
	if os.Getenv("HOST") == "" {
		baseUrl = "127.0.0.1" + baseUrl
	}
	if domain == "" {
		domain = "http://" + baseUrl
	}

	db, err := storage.NewSQLiteStorage("data.db")
	if err != nil {
		log.Fatal(err)
	}

	err = storage.RunSQLiteStorage(db, "schema.sql", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api").Subrouter()

	store := httpHandler.NewStore(db)
	store.Routes(apiRoute)

	website.Routes(router)

	log.Println("Server running on", domain)
	log.Fatal(http.ListenAndServe(baseUrl, router))
}

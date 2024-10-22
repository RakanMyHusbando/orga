package req_handler

import (
	"db-api/db_interaction"
	"log"
	"net/http"
	"strings"
)

type GetResponse struct {
	data   interface{} `json:"data"`
	status uint16      `json:"status"`
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) GetResponse {
	var props []string
	if r.URL.Query().Get("prop") != "" {
		props = strings.Split(r.URL.Query().Get("prop"), ",")
	}
	var name string = r.URL.Query().Get("name")
	var discord_id string = r.URL.Query().Get("discord_id")

	var getRes GetResponse
	jsonGet, err := db_interaction.Get(name, discord_id, props)
	if err != nil {
		log.Fatal(err)
		getRes.data = interface{}("error")
		getRes.status = 404
	} else {
		log.Println("successful user.get")
		getRes.data = jsonGet
		getRes.status = 200
	}

	return getRes
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

}

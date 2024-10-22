package req_handler

import (
	"db-api/db_interaction"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type response struct {
	// data   []map[string]interface{} `json:"data"`
	status uint16 `json:"status"`
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	var props []string
	fmt.Println("test")
	if r.URL.Query().Get("prop") != "" {
		props = strings.Split(r.URL.Query().Get("prop"), ",")
	}
	fmt.Println("test1")
	var name string = r.URL.Query().Get("name")
	var discord_id string = r.URL.Query().Get("discord_id")
	fmt.Println("test2")
	var res response = response{
		// []map[string]interface{}{},
		404,
	}
	fmt.Println("test3")
	_, err := db_interaction.Get(name, discord_id, props)
	fmt.Println("test4")
	if err != nil {
		fmt.Println("test5")
		log.Fatal(err)

	} else {
		fmt.Println("test6")
		log.Println("successful user.get")
		// res.data = jsonGet
		res.status = 200
	}
	fmt.Println("test7")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("test8")
	//specify HTTP status code
	w.WriteHeader(http.StatusOK)
	fmt.Println("test8")
	fmt.Println(res)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResponse)

}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

}

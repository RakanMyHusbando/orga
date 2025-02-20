package website

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	disgoauth "github.com/realTristan/disgoauth"
)

var (
	domain    string
	reqClient = &http.Client{}
	dc        *disgoauth.Client
)

func Init(dcClientId, dcClientSecret, domain string) {
	dc = disgoauth.Init(&disgoauth.Client{
		ClientID:     dcClientId,
		ClientSecret: dcClientSecret,
		RedirectURI:  domain + "/discord/auth/callback",
		Scopes:       []string{disgoauth.ScopeIdentify},
	})
}

func Routes(router *mux.Router) {
	router.HandleFunc("/discord", discordLoginHandler)
	router.HandleFunc("/discord/auth/callback", dcOAuthHandler)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

}

func createToken(lenght int) string {
	bytes := make([]byte, lenght)
	if _, err := rand.Read(bytes); err != nil {
		log.Println("Failed to generate session cookie: ", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	DiscordId    string `json:"discord_id"`
	SessionToken string `json:"session_coockie"`
	CSRFToken    string `json:"csrf_token"`
}

func NewUser(name, discordId, sessionToken, csrfToken string) *User {
	return &User{
		Name:         name,
		DiscordId:    discordId,
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
	}
}

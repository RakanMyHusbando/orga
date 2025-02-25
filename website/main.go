package website

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	disgoauth "github.com/realTristan/disgoauth"
)

func (web *Website) Routes(router *mux.Router) {
	discordRouter := router.PathPrefix("/discord").Subrouter()
	pagesRouter := router.PathPrefix("/").Subrouter()

	discordRouter.HandleFunc("/login", web.discordLoginHandler)
	discordRouter.HandleFunc("/auth/callback", web.discordOauth2RederectHandler)

	pagesRouter.HandleFunc("/user", web.userPageHandler)

	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./website/public"))))
}

func readIp(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func NewWebsite(storage *SessionStorage, dcClientId, dcClientSecret, domain string) (*Website, error) {
	return &Website{
		storage: storage,
		dc: disgoauth.Init(&disgoauth.Client{
			ClientID:     dcClientId,
			ClientSecret: dcClientSecret,
			RedirectURI:  domain + "/discord/auth/callback",
			Scopes:       []string{disgoauth.ScopeIdentify},
		}),
		httpClient: &http.Client{},
	}, nil
}

type Website struct {
	storage    *SessionStorage
	dc         *disgoauth.Client
	httpClient *http.Client
}

type User struct {
	DiscordId    string
	Ip           string
	SessionToken string
}

/*
If sessionToken is an int, it will be used to generate a random token.
If sessionToken is a string, it will be used as the token.
If sessionToken is any other type, an empty string will be used.
*/
func newUser(discordId, Ip string, sessionToken any) *User {
	var st string
	stType := reflect.TypeOf(sessionToken).Kind()
	if stType == reflect.Int {
		st = createToken(sessionToken.(int))
	} else if stType == reflect.String {
		st = sessionToken.(string)
	}
	return &User{
		DiscordId:    discordId,
		Ip:           Ip,
		SessionToken: st,
	}
}

func createToken(lenght int) string {
	bytes := make([]byte, lenght)
	if _, err := rand.Read(bytes); err != nil {
		log.Println("Failed to generate session cookie: ", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func (web *Website) authorize(r *http.Request) error {
	user, err := web.storage.Select(r.FormValue("discord_id"))
	sessionToken, err := r.Cookie("session_token")
	if err != nil || sessionToken.Value == "" || sessionToken.Value != user.SessionToken {
		return fmt.Errorf("Failed to authorize user")
	}
	return nil
}

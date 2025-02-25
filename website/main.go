package website

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	disgoauth "github.com/realTristan/disgoauth"
)

func (web *Website) Routes() {
	http.HandleFunc("/htmx/headline", web.headlineHandler)

	http.HandleFunc("/discord/login", web.discordLoginHandler)
	http.HandleFunc("/discord/auth/callback", web.discordOauth2RederectHandler)

	http.HandleFunc("/user", web.userPageHandler)

	http.Handle("/", http.FileServer(http.Dir("./website/public")))
}

func (web *Website) logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	}
	web.storage.Delete(sessionToken.Value)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func readIp(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return strings.Split(IPAddress, ":")[0]
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
	discordId := r.FormValue("discord_id")
	user, err := web.storage.Select(discordId)
	stCoockie, err := r.Cookie("session_token")
	fmt.Println("discord_id:", discordId, "\nsession_token:", stCoockie.Value)
	if err != nil || stCoockie.Value == "" || user.SessionToken != stCoockie.Value {
		return fmt.Errorf("Unauthorized")
	}
	return nil
}

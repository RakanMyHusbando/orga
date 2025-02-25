package website

import (
	"log"
	"net/http"
	"time"

	disgoauth "github.com/realTristan/disgoauth"
)

func (web *Website) discordLoginHandler(w http.ResponseWriter, r *http.Request) {
	web.dc.RedirectHandler(w, r, "")
}

func (web *Website) discordOauth2RederectHandler(w http.ResponseWriter, r *http.Request) {
	codeFromURLParamaters := r.URL.Query()["code"][0]
	accessToken, err := web.dc.GetOnlyAccessToken(codeFromURLParamaters)
	if err != nil {
		log.Println(err)
	}
	data, err := disgoauth.GetUserData(accessToken)
	if err != nil {
		log.Println(err)
	}
	user := newUser(data["id"].(string), readIp(r), 32)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    user.SessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	web.storage.Insert(user)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

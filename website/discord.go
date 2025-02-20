package website

import (
	"log"
	"net/http"
	"time"

	disgoauth "github.com/realTristan/disgoauth"
)

func discordLoginHandler(w http.ResponseWriter, r *http.Request) {
	dc.RedirectHandler(w, r, "")
}

func dcOAuthHandler(w http.ResponseWriter, r *http.Request) {
	codeFromURLParamaters := r.URL.Query()["code"][0]
	accessToken, err := dc.GetOnlyAccessToken(codeFromURLParamaters)
	if err != nil {
		log.Println(err)
	}

	data, err := disgoauth.GetUserData(accessToken)
	if err != nil {
		log.Println(err)
	}

	sessionToken := createToken(32)
	csrfToken := createToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	_ = data

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

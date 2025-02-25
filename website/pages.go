package website

import (
	"log"
	"net/http"
)

func (web *Website) userPageHandler(w http.ResponseWriter, r *http.Request) {
	if err := web.authorize(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, "./website/pages/user.html")
}

func (web *Website) logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	}
	web.storage.Delete(sessionToken.Value)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

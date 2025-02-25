package website

import (
	"fmt"
	"net/http"
	"text/template"
)

func loadError(page string) error {
	return fmt.Errorf("Failed to load %s.html", page)
}

func parsePage(page string) (*template.Template, error) {
	s := fmt.Sprintf("website/pages/%s.html", page)
	return template.ParseFiles(s)
}

func (web *Website) userPageHandler(w http.ResponseWriter, r *http.Request) {
	if err := web.authorize(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tmpl, err := parsePage("user")
	if err != nil {
		http.Error(w, loadError("user").Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, loadError("user").Error(), http.StatusInternalServerError)
	}
}

func (web *Website) headlineHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parsePage("headline")
	if err != nil {
		http.Error(w, loadError("headline").Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, loadError("headline").Error(), http.StatusInternalServerError)
	}
}

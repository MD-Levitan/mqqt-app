package router

import (
	"net/http"

	"github.com/MD-Levitan/mqqt-app/config"
)

func loginWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "login.html", nil)
}

func viewWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "view.html", nil)
}

func logoutWebHandler(w http.ResponseWriter, r *http.Request) {
	if err := logout(w, r); err != nil {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

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
